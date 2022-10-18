// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port           = flag.Int("port", 50151, "The server port")
	start_session  = flag.Uint64("start", 1, "The starting session ID")
	max_session    = flag.Uint64("max", 8192, "The maximum session ID")
	update         = flag.Int("update", 10, "Delay for each session update run")
)

type server struct {
	fw.UnimplementedSessionTableServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	fw.RegisterSessionTableServer(s, &server{})

	reflection.Register(s)

	init_sessionoffload()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
