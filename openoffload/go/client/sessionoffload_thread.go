// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.
//
// This file contains the logic for the client thead, which will run
// and verify sessions with the server, removing sessions which are
// closed.

package main

import (
	"context"
	"io"
        "log"
        "time"

        fw "github.com/opiproject/sessionOffload/sessionoffload/v2/gen/go"
	"google.golang.org/grpc"
)

func do_client_background(conn grpc.ClientConnInterface, ch chan<- bool) {
	log.Printf("----- Entered do_client_background -----")

	// Count sessions
	count := 0

	// Session Offload client
	client := fw.NewSessionTableClient(conn)

	for {
		count = 0

		time.Sleep(time.Duration(10) * time.Second)

		// Get all the sessions
		request_args := &fw.SessionRequestArgs{
			StartSession: 0,
		}

		stream, err := client.GetClosedSessions(context.Background(), request_args)
		if err != nil {
			log.Fatalf("open stream error %v", err)
		}

		done := make(chan bool)

		go func(client fw.SessionTableClient) {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					done <- true //means stream is finished
					return
				}
				if err != nil {
					log.Fatalf("cannot receive %v", err)
				}
				count++

				sess_id := &fw.SessionId{
					SessionId: resp.SessionId,
				}

				// Clear out this session
				_, err = client.DeleteSession(context.Background(), sess_id)
				if err != nil {
					log.Printf("Failed removing session %d", sess_id.SessionId)
				}
			}
		}(client)

		<-done //we will wait until all response is received
		log.Printf("finished")

		// Get current sessions
		all_sessions, sess_err := client.GetAllSessions(context.Background(), request_args)
		if sess_err != nil {
			log.Printf("Error contacting server")
			done <- true
			return
		}

		if len(all_sessions.SessionInfo) == 0 && count == 0 {
			// Exit as we processed all closed sessions
			log.Printf("All sessions closed, returning")
			done <- true
			return
		}
	}
}
