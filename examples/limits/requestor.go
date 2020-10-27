package main

import (
	"fmt"
	"github.com/rsms/gotalk"
	"time"
)

func requestor(port string) {
	// Connect to our "responder"
	s, err := gotalk.Connect("tcp", "localhost:"+port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("connected to %q\n", s.Addr())

	// We use a single channel for receiving all responses
	reschan := make(chan gotalk.Response)

	// In this example all requests are the same
	req := gotalk.NewRequest("echo", []byte("hello"))
	sendRequest := func() {
		err := s.SendRequest(req, reschan)
		if err != nil {
			panic(err)
		}
	}

	// Send 10 requests at the same time
	for i := 0; i != 10; i++ {
		sendRequest()
	}

	// Read all responses
	for i := 0; i != 10; i++ {
		res := <-reschan
		if res.IsRetry() {
			fmt.Printf("requestor: retrying request in %v\n", res.Wait)
			if res.Wait == 0 {
				sendRequest()
			} else {
				go func(wait time.Duration) {
					time.Sleep(res.Wait)
					sendRequest()
				}(res.Wait)
			}
			i--
		} else if res.IsError() {
			panic(fmt.Sprintf("error response: %v", res.Error()))
		} else {
			fmt.Printf("requestor: received response: %q\n", string(res.Data))
		}
	}

	fmt.Printf("done. Closing socket\n")
	s.Close()
}
