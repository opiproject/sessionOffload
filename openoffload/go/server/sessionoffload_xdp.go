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
		session.Ipv4.Sport = uint16(req.SourcePort)
		session.Ipv4.Dport = uint16(req.DestinationPort)
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
		session.Ipv6.Sport = uint16(req.SourcePort)
		session.Ipv6.Dport = uint16(req.DestinationPort)
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
	var s fw.SessionResponse

	s.SessionId = uint64(session_id)
	s.InPackets = entry.In_packets
	s.OutPackets = entry.Out_packets
	s.InBytes = entry.In_bytes
	s.OutBytes = entry.Out_bytes
	s.SessionState = fw.SessionState(entry.Session_state)
	s.SessionCloseCode = fw.SessionCloseCode(entry.Session_close_code)

	return &s, nil
}

func xdp_update_map_entry(s *session)(error) {
	var entry xdp_session

	mapdata, err := sessionmap.Lookup(s.session_id)
	if err != nil {
		log.Printf("Error reading session map")
		return err
	}
	buf := bytes.NewBuffer(mapdata)

	if err := binary.Read(buf, binary.LittleEndian, &entry); err != nil {
		log.Printf("Error reading data from session map")
		return err
	}

	// Update the entry
	entry.In_lif = uint32(s.in_lif)
	entry.Out_lif = uint32(s.out_lif)
	if s.ip_version == fw.IpVersion__IPV4 {
		entry.Ip_version = 0
		entry.Ipv4.Saddr = s.source_ip
		entry.Ipv4.Daddr = s.destination_ip
		entry.Ipv4.Sport = uint16(s.source_port)
		entry.Ipv4.Dport = uint16(s.destination_port)
		entry.Action.Action_next_hop = s.action.action_next_hop
	} else {
		entry.Ip_version = 1
		rbuf := bytes.NewReader(s.source_ipv6)
		err := binary.Read(rbuf, binary.LittleEndian, &entry.Ipv6.Saddr)
		if err != nil {
			log.Printf("Failed reading IPv6 source")
			return err
		}
		rbuf = bytes.NewReader(s.destination_ipv6)
		err = binary.Read(rbuf, binary.LittleEndian, &entry.Ipv6.Daddr)
		if err != nil {
			log.Printf("Failed reading IPv6 destination")
			return err
		}
		entry.Ipv6.Sport = uint16(s.source_port)
		entry.Ipv6.Dport = uint16(s.destination_port)
		rbuf = bytes.NewReader(s.action.action_next_hop_v6)
		err = binary.Read(rbuf, binary.LittleEndian, &entry.Action.Action_next_hop_v6)
		if err != nil {
			log.Printf("Error reading V6 next hop")
			return err
		}
	}
	entry.In_packets = s.in_packets
	entry.Out_packets = s.out_packets
	entry.In_bytes = s.in_bytes
	entry.Out_bytes = s.out_bytes
	entry.Session_state = uint32(s.session_state)
	entry.Session_close_code = uint32(s.session_close_code)

	buf2 := &bytes.Buffer{}

	err = binary.Write(buf2, binary.LittleEndian, entry)
	if err != nil {
		log.Printf("Failed writing session map for ID %d: %v", s.session_id, err)
	}

	err = sessionmap.Upsert(s.session_id, buf2.Bytes())
	if err != nil {
		log.Printf("Failed adding session %d into eBPF map: %v", s.session_id, err)
	}
	return nil
}

func xdp_add_session(sessionId uint32, req *fw.SessionRequest) error {
	entry, err := xdp_convert_session(req)
	if err != nil {
		log.Printf("Failed converting session map structure")
		return err
	}

	log.Printf("xdp_add_session: Adding ID [%d] [%v]", sessionId, entry)

	buf := &bytes.Buffer{}

	err = binary.Write(buf, binary.LittleEndian, entry)
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

	resp, err2 := xdp_convert_session_to_proto(session, entry)
	if err2 != nil {
		log.Printf("Failed converting map entry to proto struct")
		return nil, err2
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
