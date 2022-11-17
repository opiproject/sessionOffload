// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.
//
// This file contains the logic for the session_update thread. This is meant to
// simulate the random updating of sessions in the session table and would not
// be used in a setup on actual hardware.
//

package main

import (
	"bytes"
	"encoding/binary"
        "log"

	"github.com/dropbox/goebpf"
        fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
)

type ipv4_t struct {
	Saddr uint32
	Daddr uint32
	Sport uint16
	Dport uint16
}

type ipv6_t struct {
	Saddr [4]uint32
	Daddr [4]uint32
	Sport uint16
	Dport uint16
}

type action_t struct {
	Action             uint32
	Action_next_hop    uint32
	Action_next_hop_v6 [4]uint32
}

// The Session Table struct
type xdp_session struct {
	In_lif             uint32
	Out_lif            uint32
	Ip_version         int32
	Ipv4               ipv4_t
	Ipv6               ipv6_t
	Protocol           uint32
	Action             action_t
	In_packets         uint64
	Out_packets        uint64
	In_bytes           uint64
	Out_bytes          uint64
	Session_state      uint32
	Session_close_code uint32
	/* Start Time: start_time         time.Time */
	/* End Time: end_time           time.Time */
}

// The path to the session table map
var map_path = "/sys/fs/bpf/session_table"

// The eBPF session map pointer
var (
	sessionmap *goebpf.EbpfMap
)

func init_xdp() {
	var err error

	sessionmap, err = goebpf.NewMapFromExistingMapByPath(map_path)
	if err != nil {
		panic(err)
	}
	log.Printf("Found map %s of type %s\n", sessionmap.Name, sessionmap.Type)
}

func xdp_convert_session(req *fw.SessionRequest)(*xdp_session, error) {
	var session xdp_session

	session.In_lif = uint32(req.InLif)
	session.Out_lif = uint32(req.OutLif)
	session.Action.Action = uint32(req.Action.ActionType.Number())
	if req.IpVersion == fw.IpVersion__IPV4 {
		session.Ip_version = 0 // .FIXME: Hardcoded
		session.Ipv4.Saddr = req.SourceIp
		session.Ipv4.Daddr = req.DestinationIp
		session.Ipv4.Sport = uint16(req.SourcePort >> 16)
		session.Ipv4.Dport = uint16(req.DestinationPort >> 16)
		session.Action.Action_next_hop = req.Action.ActionNextHop
	} else {
		session.Ip_version = 1 // .FIXME: Hardcoded
		rbuf := bytes.NewReader(req.SourceIpv6)
		err := binary.Read(rbuf, binary.LittleEndian, &session.Ipv6.Saddr)
		if err != nil {
			log.Printf("Failed reading IPv6 source")
			return nil, err
		}
		rbuf = bytes.NewReader(req.DestinationIpv6)
		err = binary.Read(rbuf, binary.LittleEndian, &session.Ipv6.Daddr)
		if err != nil {
			log.Printf("Failed reading IPv6 destination")
			return nil, err
		}
		session.Ipv6.Sport = uint16(req.SourcePort >> 16)
		session.Ipv6.Dport = uint16(req.DestinationPort >> 16)
		rbuf = bytes.NewReader(req.Action.ActionNextHopV6)
		err = binary.Read(rbuf, binary.LittleEndian, &session.Action.Action_next_hop_v6)
		if err != nil {
			log.Printf("Error reading V6 next hop")
			return nil, err
		}
	}
	session.Protocol = uint32(req.ProtocolId.Number())
	session.In_packets = 0
	session.Out_packets = 0
	session.In_bytes = 0
	session.Out_bytes = 0
	session.Session_state = 0
	session.Session_close_code = 0

	return &session, nil
}

func xdp_convert_session_to_proto(session_id uint32, entry xdp_session)(*fw.SessionResponse, error) {
	var session fw.SessionResponse

	session.SessionId = uint64(session_id)
	session.InPackets = entry.In_packets
	session.OutPackets = entry.Out_packets
	session.InBytes = entry.In_bytes
	session.OutBytes = entry.Out_bytes
	session.SessionState = fw.SessionState(entry.Session_state)
	session.SessionCloseCode = fw.SessionCloseCode(entry.Session_close_code)

	return &session, nil
}

func xdp_add_session(sessionId uint32, req *fw.SessionRequest) error {
	session, err := xdp_convert_session(req)
	if err != nil {
		log.Printf("Failed converting session map structure")
		return err
	}

	buf := &bytes.Buffer{}

	err = binary.Write(buf, binary.LittleEndian, session)
	if err != nil {
		log.Printf("Failed writing session map for ID %d: %v", sessionId, err)
	}

	err = sessionmap.Upsert(sessionId, buf.Bytes())
	if err != nil {
		log.Printf("Failed adding session %d into eBPF map: %v", sessionId, err)
	}

	return err
}

func xdp_get_session(session uint32)(*fw.SessionResponse, error) {
	var resp *fw.SessionResponse
	var entry xdp_session

	mapdata, err := sessionmap.Lookup(session)
	if err != nil {
		log.Printf("Error reading session map")
		return nil, err
	}
	buf := bytes.NewBuffer(mapdata)

	if err := binary.Read(buf, binary.LittleEndian, &entry); err != nil {
		log.Printf("Error reading data from session map")
		return nil, err
	}

	resp, err = xdp_convert_session_to_proto(session, entry)
	if err != nil {
		log.Printf("Failed converting map entry to proto struct")
		return nil, err
	}

	return resp, nil
}

func xdp_del_session(session uint32)(*fw.SessionResponse, error) {
	session_entry, err := xdp_get_session(session)
	if err != nil {
		log.Printf("Cannot find session %d to delete", session)
		return nil, err
	}

	err = sessionmap.Delete(session)
	if err != nil {
		log.Printf("Error deleting session for key %d", session)
		return nil, err
	}

	return session_entry, nil
}
