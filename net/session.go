package net

import (
	"io"
	"net"
)

type Session interface {
	ID() int
	Addr() net.Addr
	io.Writer
	io.Reader
}

func newSession(id int, conn net.Conn) Session {
	return &session{id: id, conn: conn}
}

type session struct {
	id   int
	conn net.Conn
}

func (s session) ID() int {
	return s.id
}

func (s session) Addr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *session) Write(p []byte) (n int, err error) {
	return s.conn.Write(p)
}

func (s *session) Read(p []byte) (n int, err error) {
	return s.conn.Read(p)
}
