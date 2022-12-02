// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.

package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type action_params struct {
	action_type        fw.ActionType
	action_next_hop    uint32
	action_next_hop_v6 []byte
}

// Per session state in the table
type session struct {
	session_id         uint32
	in_lif             int32
	out_lif            int32
	ip_version         fw.IpVersion
	source_ip          uint32
	source_ipv6        []byte
	source_port        uint32
	destination_ip     uint32
	destination_ipv6   []byte
	destination_port   uint32
	protocol_id        fw.ProtocolId
	action             action_params
	cache_timeout      uint32
	in_packets         uint64
	out_packets        uint64
	in_bytes           uint64
	out_bytes          uint64
	session_state      fw.SessionState
	session_close_code fw.SessionCloseCode
	request_status     fw.RequestStatus
	start_time         time.Time
	end_time           time.Time
}

// Mutex for global session table
var session_lock sync.RWMutex

// The global session table
var sessions map[uint32]session

// The last session ID we allocated
var last uint32

// The max session ID we can allocate
var max uint32

func find_session_by_7_tuple(in_lif, out_lif int32, source_ipv4 uint32, source_ipv6 []byte,
			dest_ipv4 uint32, dest_ipv6 []byte,
			src_port, dst_port uint32,
			protocol_id *fw.ProtocolId, ip_version *fw.IpVersion) uint32 {
	for _, v := range sessions {
		session_lock.RLock()

		if *ip_version == fw.IpVersion__IPV4 {
			if v.in_lif == in_lif &&
			   v.out_lif == out_lif &&
			   v.source_ip == source_ipv4 &&
			   v.destination_ip == dest_ipv4 &&
			   v.source_port == src_port &&
			   v.destination_port == dst_port &&
			   v.protocol_id == *protocol_id {
				session_lock.RUnlock()
				return v.session_id
			}
		} else if *ip_version == fw.IpVersion__IPV6 {
			srcv6 := bytes.Compare(v.source_ipv6, source_ipv6)
			dstv6 := bytes.Compare(v.destination_ipv6, dest_ipv6)

			if v.in_lif == in_lif &&
			   v.out_lif == out_lif &&
			   srcv6 == 0 &&
			   dstv6 == 0 &&
			   v.source_port == src_port &&
			   v.destination_port == dst_port &&
			   v.protocol_id == *protocol_id {
				session_lock.RUnlock()
				return v.session_id
			}
		}

		session_lock.RUnlock()
	}

	return 0
}

func init_sessionoffload() {
	rand.Seed(time.Now().Unix())
	sessions = make(map[uint32]session)
	last = uint32(*start_session)
	max  = uint32(*max_session)
	if *update != 0 {
		go session_update()
	}
}

func next_session_id() (uint32, error) {
	var cnt uint32

	session_lock.RLock()
	defer session_lock.RUnlock()

	cnt = 0
	for {
		cnt += 1
		if cnt == max {
			return 0, errors.New("Session table is full")
		}

		if last == max {
			last = 1
		}

		if _, found := sessions[cnt]; found {
			last += 1
		} else {
			return last, nil
		}
	}
}

func (s *server) AddSession(stream fw.SessionTable_AddSessionServer) error {
	var total int32
	var resp fw.AddSessionResponse
	var add_new_session bool

	for {
		add_new_session = true
		sr, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&resp)
		}
		if err != nil {
			return err
		}

		// Check if this session already exists
		if existing_session := find_session_by_7_tuple(sr.InLif, sr.OutLif,
								sr.SourceIp, sr.SourceIpv6,
								sr.DestinationIp, sr.DestinationIpv6,
								sr.SourcePort, sr.DestinationPort,
								&sr.ProtocolId, &sr.IpVersion); existing_session != 0 {
			log.Printf("Existing session with 7-tuple found")
			session_resp_err := fw.SessionResponseError{
				SessionId: uint64(existing_session),
				ErrorStatus: fw.RequestStatus_value["_REJECTED_SESSION_ALREADY_EXISTS"],
			}
			resp.ResponseError = append(resp.ResponseError, &session_resp_err)
			add_new_session = false
		}

		newSessionId, err := next_session_id()
		if err != nil {
			log.Printf("Error getting new session ID: %v", err)
			session_resp_err := fw.SessionResponseError{
				ErrorStatus: fw.RequestStatus_value["_REJECTED_SESSION_TABLE_FULL"],
			}
			resp.ResponseError = append(resp.ResponseError, &session_resp_err)
			add_new_session = false
		}

		if add_new_session {
			new_session := session {
				session_id:       newSessionId,
				in_lif:           sr.InLif,
				out_lif:          sr.OutLif,
				ip_version:       sr.IpVersion,
				source_ip:        sr.SourceIp,
				source_ipv6:      sr.SourceIpv6,
				source_port:      sr.SourcePort,
				destination_ip:   sr.DestinationIp,
				destination_ipv6: sr.DestinationIpv6,
				destination_port: sr.DestinationPort,
				protocol_id:      sr.ProtocolId,
				action:           action_params{
					action_type:        sr.Action.ActionType,
					action_next_hop:    sr.Action.ActionNextHop,
					action_next_hop_v6: sr.Action.ActionNextHopV6,
				},
				cache_timeout:    sr.CacheTimeout,
				in_packets:       0,
				out_packets:      0,
				in_bytes:         0,
				out_bytes:        0,
				start_time:       time.Now(),
			}

			session_lock.Lock()
			sessions[newSessionId] = new_session
			session_lock.Unlock()

			if *xdp_backend {
				err = xdp_add_session(newSessionId, sr)
				if err != nil {
					log.Printf("Failed adding session %d into eBPF map", newSessionId)
				}
			}

			total++
		}

		log.Printf("%+v\n", sr)
	}
}

func (s *server) GetSession(ctx context.Context, in *fw.SessionId) (*fw.SessionResponse, error) {
	session_lock.RLock()
	defer session_lock.RUnlock()

	session, valid := sessions[uint32(in.SessionId)]
	if !valid {
		log.Printf("Session not found")
		return nil, errors.New("Session not found")
	}

	if *xdp_backend {
		resp, err := xdp_get_session(uint32(in.SessionId))
		if err != nil {
			log.Printf("Failed adding session %d into eBPF map", in.SessionId)
		}
		return resp, nil
	} else {
		return &fw.SessionResponse{
			SessionId:        uint64(session.session_id),
			InPackets:        session.in_packets,
			OutPackets:       session.out_packets,
			InBytes:          session.in_bytes,
			OutBytes:         session.out_bytes,
			SessionState:     session.session_state,
			SessionCloseCode: session.session_close_code,
			RequestStatus:    session.request_status,
			StartTime:        timestamppb.New(session.start_time),
			EndTime:          timestamppb.New(session.end_time),
		}, nil
	}
}

func (s *server) DeleteSession(ctx context.Context, in *fw.SessionId) (*fw.SessionResponse, error) {
	session_lock.Lock()
	defer session_lock.Unlock()

	log.Printf("----- DELETE SESSION -----")
	log.Printf("Looking for session %d", in.SessionId)

	session, valid := sessions[uint32(in.SessionId)]
	if !valid {
		log.Printf("Session not found")
		return nil, errors.New("Session not found")
	}

	return_session := &fw.SessionResponse{
		SessionId:        uint64(session.session_id),
		InPackets:        session.in_packets,
		OutPackets:       session.out_packets,
		InBytes:          session.in_bytes,
		OutBytes:         session.out_bytes,
		SessionState:     session.session_state,
		SessionCloseCode: session.session_close_code,
		RequestStatus:    session.request_status,
		StartTime:        timestamppb.New(session.start_time),
		EndTime:          timestamppb.New(session.end_time),
	}

	delete(sessions, uint32(in.SessionId))

	if *xdp_backend {
		_, err := xdp_del_session(uint32(in.SessionId))
		if err != nil {
			log.Printf("Failed adding session %d into eBPF map", in.SessionId)
		}
	}

	return return_session, nil
}

func (s *server) GetAllSessions(ctx context.Context, in *fw.SessionRequestArgs) (*fw.SessionResponses, error) {
	var return_sessions fw.SessionResponses

	log.Printf("----- GET ALL SESSIONS -----")
	log.Printf("Starting with session %d", in.StartSession)

	session_lock.RLock()
	defer session_lock.RUnlock()

	for k, v := range sessions {
		// Skip if requested session start is greater than current session
		if k < uint32(in.StartSession) {
			continue
		}

		// Collect the response
		if *xdp_backend {
			resp, err := xdp_get_session(k)
			if err != nil {
				log.Printf("Failed adding session %d into eBPF map", k)
			}

			return_sessions.SessionInfo = append(return_sessions.SessionInfo, resp)
		} else {
			return_session := &fw.SessionResponse{
				SessionId:        uint64(v.session_id),
				InPackets:        v.in_packets,
				OutPackets:       v.out_packets,
				InBytes:          v.in_bytes,
				OutBytes:         v.out_bytes,
				SessionState:     v.session_state,
				SessionCloseCode: v.session_close_code,
				RequestStatus:    v.request_status,
				StartTime:        timestamppb.New(v.start_time),
				EndTime:          timestamppb.New(v.end_time),
			}

			return_sessions.SessionInfo = append(return_sessions.SessionInfo, return_session)
		}
	}

	return &return_sessions, nil
}

func (s *server) GetClosedSessions(in *fw.SessionRequestArgs, stream fw.SessionTable_GetClosedSessionsServer) error {
	// Use a wait group to allow for process concurrency
	var wg sync.WaitGroup

	session_lock.RLock()
	defer session_lock.RUnlock()

	for k, v := range sessions {
		if *xdp_backend {
			resp, err := xdp_get_session(k)
			if err != nil {
				log.Printf("Failed adding session %d into eBPF map", k)
			}

			if resp.SessionState != fw.SessionState__CLOSED {
				continue
			}

			wg.Add(1)

			go func(session *fw.SessionResponse) {
				defer wg.Done()

				if err := stream.Send(resp); err != nil {
					log.Printf("Error sending closed session reply: %v", err)
				}
			}(resp)
		} else {
			if v.session_state != fw.SessionState__CLOSED {
				continue
			}

			// Stream this session
			wg.Add(1)

			go func(v session) {
				defer wg.Done()

				// Collect the response
				return_session := &fw.SessionResponse{
					SessionId:        uint64(v.session_id),
					InPackets:        v.in_packets,
					OutPackets:       v.out_packets,
					InBytes:          v.in_bytes,
					OutBytes:         v.out_bytes,
					SessionState:     v.session_state,
					SessionCloseCode: v.session_close_code,
					RequestStatus:    v.request_status,
					StartTime:        timestamppb.New(v.start_time),
					EndTime:          timestamppb.New(v.end_time),
				}

				if err := stream.Send(return_session); err != nil {
					log.Printf("send error %v", err)
				}
			}(v)
		}
	}

	wg.Wait()
	return nil
}
