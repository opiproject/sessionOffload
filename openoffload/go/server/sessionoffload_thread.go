// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.
//
// This file contains the logic for the session_update thread. This is meant to
// simulate the random updating of sessions in the session table and would not
// be used in a setup on actual hardware.
//

package main

import (
	"log"
	"math/rand"
	"time"

	fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
)

// Update packet counters
func session_update_packet_counters(v *session) {
	// Increment packet counters
	v.in_packets  += uint64(rand.Intn(50))
	v.out_packets += uint64(rand.Intn(700))
	v.in_bytes    += uint64(rand.Intn(7000))
	v.out_bytes   += uint64(rand.Intn(500000))

	if *xdp_backend {
		if err := xdp_update_map_entry(v); err != nil {
			log.Printf("Failed updating stats for eBPF map %d", v.session_id)
		}
	}
}

func session_timeout(v *session) {
	v.session_state      = fw.SessionState__CLOSED
	v.session_close_code = fw.SessionCloseCode__TIMEOUT
	v.end_time           = time.Now()

	if *xdp_backend {
		if err := xdp_update_map_entry(v); err != nil {
			log.Printf("Failed updating session timeout for eBPF map %d", v.session_id)
		}
	}
}

func session_fin(v *session) {
	v.session_state      = fw.SessionState__CLOSED
	v.session_close_code = fw.SessionCloseCode__FINACK
	v.end_time           = time.Now()

	if *xdp_backend {
		if err := xdp_update_map_entry(v); err != nil {
			log.Printf("Failed updating session fin for eBPF map %d", v.session_id)
		}
	}
}

// Function which runs in the background updating the session table entries
func session_update() {
	states := [18]string  {
		"noop", "noop", "noop", "noop", "noop", "noop", "noop", "noop", "noop", "noop",
		"stats", "stats", "stats", "stats", "stats",
		"timeout", "clientfin", "serverfin",
	}

	for {
		log.Printf("----- session_update running -----")
		time.Sleep(time.Duration(*update) * time.Second)


		session_lock.RLock()
		for k, v := range sessions {

			if v.session_state == fw.SessionState__CLOSED {
				// Skip this session
				log.Printf("Skipping update for session %d", v.session_id)
				continue
			}

			// Get a random state
			switch state := states[rand.Intn(len(states))]; state {
			case "stats":
				log.Printf("Updating stats for session %d", v.session_id)
				session_update_packet_counters(&v)
			case "timeout":
				log.Printf("Timing out session %d", v.session_id)
				session_timeout(&v)
			case "clientfin":
				log.Printf("Session %d: CLIENTFIN", v.session_id)
				session_fin(&v)
			case "serverfin":
				log.Printf("Session %d: SERVERFIN", v.session_id)
				session_fin(&v)
			case "noop":
			default:
				log.Printf("default for session %d", v.session_id)
			}

			// Save the new session in the session map
			sessions[k] = v

			// Dump the session
			log.Printf("Session %d: ID: [%d] State: [%s] In packets/bytes [%d/%d] Out packets/bytes [%d/%d]",
				k, v.session_id, v.session_state.String(),
				v.in_packets, v.in_bytes,
				v.out_packets, v.out_bytes)

		}
		session_lock.RUnlock()
	}
}
