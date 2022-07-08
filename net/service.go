package net

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

func newTCPListener(port int) (ln net.Listener, err error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

// Start creates and serves the socket
func Start(wg *sync.WaitGroup, ctx context.Context, port int) {
	// Create a listener on the given port
	ln, err := newTCPListener(port)
	if err != nil {
		log.Fatal("Could not start TCP server:", err)
		return
	}

	go func() {
		<-ctx.Done() // Block until cancel signal
		ln.Close()   // Close listener
		log.Println("Socket stopped")
	}()

	go accept(ln, wg)
}

func accept(ln net.Listener, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement parent wait group for listener

	for {
		conn, err := ln.Accept() // Block until new connection arrives and accept
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				// Listener closed. Break out of accept loop
				break
			}

			log.Println("Error accepting client:", err)
			continue
		}

		log.Printf("New connection (%s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		buf := make([]byte, 2048)

		_, err := conn.Read(buf)
		if err != nil {
			if op, ok := err.(*net.OpError); ok {
				if op.Op == "read" {
					log.Printf("Client closed connection (%s)", conn.RemoteAddr())
				}
			} else {
				log.Printf("Closing connection (%s): %s", conn.RemoteAddr(), err)
			}
			break
		}
	}
}
