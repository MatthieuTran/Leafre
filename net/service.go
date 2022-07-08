package net

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

// newListen creates and returns a tcp listener on the given port
func newListener(port int) (ln net.Listener, err error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

// Start creates and serves the socket
func Start(wg *sync.WaitGroup, ctx context.Context, port int) {
	// Create a listener on the given port
	ln, err := newListener(port)
	if err != nil {
		log.Fatal("Could not start TCP server:", err)
		return
	}

	log.Println("Socket started on", port)

	go func() {
		<-ctx.Done() // Block until cancel signal
		ln.Close()   // Close listener
		log.Println("Socket stopped")
	}()

	go accept(ln, wg)
}

// accept blocks until a new connection arrives and accepts the connection
func accept(ln net.Listener, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement parent wait group for listener

	for {
		conn, err := ln.Accept() // Block until new connection and accept
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				// Listener closed. Break out of accept loop
				break
			}

			log.Println("Error accepting client:", err)
			continue
		}

		go HandleConn(conn)
	}
}
