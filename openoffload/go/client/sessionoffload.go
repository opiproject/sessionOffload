// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.

package main

import (
	"context"
	"log"

	fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
	"google.golang.org/grpc"
)

func do_sessionoffload(conn grpc.ClientConnInterface, ctx context.Context) {
	log.Printf("Entered do_sessionoffload")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Session Offload client
	client := fw.NewSessionTableClient(conn)

	// Load some sessions
	add_session_requests := []*fw.SessionRequest{
		{InLif: 1, OutLif: 4, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x10101001, DestinationPort: 443, ProtocolId: fw.ProtocolId__TCP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
		{InLif: 2, OutLif: 4, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x10101010, DestinationPort: 443, ProtocolId: fw.ProtocolId__TCP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
		{InLif: 3, OutLif: 4, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x10101011, DestinationPort: 443, ProtocolId: fw.ProtocolId__TCP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
	}

	stream, err := client.AddSession(ctx)
	if err != nil {
		log.Printf("create stream: %v", err)
	}

	for _, session := range add_session_requests {
		if err := stream.Send(session); err != nil {
			log.Printf("Failed sending session")
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("close and receive: %v", err)
	}

	log.Printf("%+v\n", response)

	// Load a few more sessions
	add_session_requests2 := []*fw.SessionRequest{
		{InLif: 9, OutLif: 5, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x101a1a01, DestinationPort: 800, ProtocolId: fw.ProtocolId__UDP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
		{InLif: 8, OutLif: 6, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x101a1a10, DestinationPort: 800, ProtocolId: fw.ProtocolId__UDP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
		{InLif: 7, OutLif: 6, IpVersion: fw.IpVersion__IPV4, SourceIp: 0x101a1a11, DestinationPort: 800, ProtocolId: fw.ProtocolId__UDP, Action: &fw.ActionParameters{ActionType: fw.ActionType__DROP}, },
	}

	stream, err = client.AddSession(ctx)
	if err != nil {
		log.Printf("create stream: %v", err)
	}

	for _, session := range add_session_requests2 {
		if err := stream.Send(session); err != nil {
			log.Printf("Failed sending session")
		}
	}

	response, err = stream.CloseAndRecv()
	if err != nil {
		log.Printf("close and receive: %v", err)
	}

	log.Printf("%+v\n", response)
}
