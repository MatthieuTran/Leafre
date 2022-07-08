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
func newListener(host string, port int) (ln net.Listener, err error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
}

// Start creates and serves the socket
//
// A SocketOption should at least have a port supplied
func Start(wg *sync.WaitGroup, ctx context.Context) func(host string, port int) {
	return func(host string, port int) {
		// Create a listener on the given port
		ln, err := newListener(host, port)
		if err != nil {
			log.Fatal("Could not start TCP server:", err)
			return
		}

		log.Printf("Socket started on %s:%d", host, port)

		go func() {
			<-ctx.Done() // Block until cancel signal
			ln.Close()   // Close listener
			log.Println("Socket stopped")
		}()

		go accept(wg)(ln)
	}
}

// accept blocks until a new connection arrives and accepts the connection
func accept(wg *sync.WaitGroup) func(ln net.Listener) {
	return func(ln net.Listener) {
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
}
