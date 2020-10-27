package main

import (
	"fmt"
	"github.com/rsms/gotalk"
	"time"
)

func responder(port string) {
	// A simple echo operation with a 500ms response delay
	gotalk.HandleBufferRequest("echo", func(s *gotalk.Sock, op string, buf []byte) ([]byte, error) {
		fmt.Printf("responder: handling request\n")
		return buf, nil
	})

	// Start a server
	s, err := gotalk.Listen("tcp", "localhost:"+port)
	if err != nil {
		panic(err)
	}

	// Print when we receive heartbeats
	s.OnHeartbeat = func(load int, t time.Time) {
		fmt.Printf("responder: received heartbeat: load=%v, time=%v\n", load, t)
	}

	// Configure limits with a read timeout of 200 milliseconds
	s.Limits = &gotalk.Limits{
		BufferRequests: gotalk.Unlimited,
		ReadTimeout:    100 * time.Millisecond,
	}

	// Accept connections
	fmt.Printf("responder: listening at %q\n", s.Addr())
	go s.Accept()
}
