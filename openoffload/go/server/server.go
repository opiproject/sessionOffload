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
	start_session  = flag.Uint("start", 1, "The starting session ID")
	max_session    = flag.Uint("max", 8192, "The maximum session ID")
	update         = flag.Int("simulate", 0, "Enable simulation with a delay per run, disabled by default")
	xdp_backend    = flag.Bool("backend", false, "Enable the XDP backend")
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

	if *xdp_backend {
		init_xdp()
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
