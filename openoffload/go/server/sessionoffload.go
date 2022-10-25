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
	session_id         uint64
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
var sessions map[uint64]session

// The last session ID we allocated
var last uint64

// The max session ID we can allocate
var max uint64

// Function which runs in the background updating the session table entries
func session_update() {
	for {
		log.Printf("----- session_update running -----")
		time.Sleep(time.Duration(*update) * time.Second)


		for k, v := range sessions {
			session_lock.RLock()

			// Increment packet counters
			v.in_packets  += uint64(rand.Intn(1000))
			v.out_packets += uint64(rand.Intn(1000))
			v.in_bytes    += uint64(rand.Intn(100000))
			v.out_bytes   += uint64(rand.Intn(100000))

			// Save the new session in the session map
			sessions[k] = v

			// Use v for printing the output again
			v = sessions[k]

			// Dump the session
			log.Printf("Session %d: ID: [%d] State: [%s] In packets/bytes [%d/%d] Out packets/bytes [%d/%d]",
				k, v.session_id, v.session_state.String(),
				v.in_packets, v.in_bytes,
				v.out_packets, v.out_bytes)

			session_lock.RUnlock()
		}
	}
}

func find_session_by_7_tuple(in_lif, out_lif int32, source_ipv4 uint32, source_ipv6 []byte,
			dest_ipv4 uint32, dest_ipv6 []byte,
			src_port, dst_port uint32,
			protocol_id *fw.ProtocolId, ip_version *fw.IpVersion) uint64 {
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
	sessions = make(map[uint64]session)
	last = *start_session
	max  = *max_session
	go session_update()
}

func next_session_id() (uint64, error) {
	var cnt uint64

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
				SessionId: existing_session,
				ErrorStatus: fw.RequestStatus_value["_REJECTED_SESSION_ALREADY_EXISTS"],
			}
			resp.ResponseError = append(resp.ResponseError, &session_resp_err)
			add_new_session = false
		}

		newSessionId, err := next_session_id()
		if err != nil {
			log.Printf("Error getting new session ID: %v", err)
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
			}

			session_lock.Lock()
			sessions[newSessionId] = new_session
			session_lock.Unlock()

			total++
		}

		log.Printf("%+v\n", sr)
	}
}

func (s *server) GetSession(ctx context.Context, in *fw.SessionId) (*fw.SessionResponse, error) {
	session_lock.RLock()
	defer session_lock.RUnlock()

	session, valid := sessions[in.SessionId]
	if !valid {
		log.Printf("Session not found")
		return nil, errors.New("Session not found")
	}

	return &fw.SessionResponse{
		SessionId:        session.session_id,
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

func (s *server) DeleteSession(ctx context.Context, in *fw.SessionId) (*fw.SessionResponse, error) {
	session_lock.Lock()
	defer session_lock.Unlock()

	session, valid := sessions[in.SessionId]
	if !valid {
		log.Printf("Session not found")
		return nil, errors.New("Session not found")
	}

	return_session := &fw.SessionResponse{
		SessionId:        session.session_id,
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

	delete(sessions, in.SessionId)

	return return_session, nil
}

func (s *server) GetAllSession(ctx context.Context, in *fw.SessionRequestArgs) (*fw.SessionResponses, error) {
	var return_sessions fw.SessionResponses

	session_lock.RLock()
	defer session_lock.RUnlock()

	for k, v := range sessions {
		// Skip if requested session start is greater than current session
		if k > in.StartSession {
			continue
		}

		// Collect the response
		return_session := &fw.SessionResponse{
			SessionId:        v.session_id,
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

	return &return_sessions, nil
}

func (s *server) GetClosedSessions(in *fw.SessionRequestArgs, stream fw.SessionTable_GetClosedSessionsServer) error {
	// Use a wait group to allow for process concurrency
	var wg sync.WaitGroup

	session_lock.RLock()
	defer session_lock.RUnlock()

	for _, v := range sessions {
		if v.session_state == fw.SessionState__CLOSED {
			continue
		}

		// Stream this session
		wg.Add(1)

		// Collect the response
		return_session := &fw.SessionResponse{
			SessionId:        v.session_id,
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
		go func(session *fw.SessionResponse) {
			defer wg.Done()

			if err := stream.Send(session); err != nil {
				log.Printf("send error %v", err)
			}
		}(return_session)
	}

	wg.Wait()
	return nil
}
