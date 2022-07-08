package tcp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	onConnected    func(net.Conn)
	onPacket       func(net.Conn, []byte)
	onDisconnected func(net.Conn, error)
}

type ServerOption func(*Server)

func NewServer(options ...ServerOption) *Server {
	s := &Server{}

	for _, option := range options {
		option(s)
	}

	return s
}

// WithOnConnected provides a onConnect callback (a function takes a `net.Conn` object) to the Server
// The callback is called on the event a client first connects to the server
func (s *Server) WithOnConnected(onConnected func(net.Conn)) *Server {
	s.onConnected = onConnected

	return s
}

// WithOnPacket provides a onPacket callback (a function takes a `net.Conn` object and a []byte slice) to the Server
// The callback is called on the event a client sends a packet to the server
func (s *Server) WithOnPacket(onPacket func(net.Conn, []byte)) *Server {
	s.onPacket = onPacket

	return s
}

// WithOnDisconnected provides a onDisconnected callback (a function takes a `net.Conn` object and an error) to the Server
// The callback is called on the event a client disconnects from the server
func (s *Server) WithOnDisconnected(onDisconnected func(net.Conn, error)) *Server {
	s.onDisconnected = onDisconnected

	return s
}

// newListen creates and returns a tcp listener on the given port
func newListener(host string, port int) (ln net.Listener, err error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
}

// Start creates and serves the socket
//
// A SocketOption should at least have a port supplied
func (s *Server) Start(wg *sync.WaitGroup, ctx context.Context) func(host string, port int) error {
	return func(host string, port int) (err error) {
		// Create a listener on the given port
		ln, err := newListener(host, port)
		if err != nil {
			return
		}

		go func() {
			<-ctx.Done() // Block until cancel signal
			ln.Close()   // Close listener
		}()

		go s.accept(wg)(ln)
		return
	}
}

// accept blocks until a new connection arrives and accepts the connection
func (s *Server) accept(wg *sync.WaitGroup) func(ln net.Listener) {
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

			go s.handleConn(conn)
		}
	}
}
