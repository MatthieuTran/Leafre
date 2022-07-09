package session

import (
	"net"
)

type Codec func(d []byte) []byte

type Session interface {
	ID() string
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Encrypt(d []byte) []byte
	Decrypt(d []byte) []byte
}

func NewSession(conn net.Conn, encrypter, decrypter Codec) Session {
	return &session{id: conn.RemoteAddr().String(), conn: conn, encrypt: encrypter, decrypt: decrypter}
}

type session struct {
	id      string
	conn    net.Conn
	encrypt Codec
	decrypt Codec
}

func (s session) ID() string {
	return s.conn.RemoteAddr().String()
}

func (s *session) Write(p []byte) (n int, err error) {
	return s.conn.Write(s.Encrypt(p))
}

func (s *session) Read(p []byte) (n int, err error) {
	return s.conn.Read(s.Decrypt(p))
}

func (s *session) Encrypt(d []byte) []byte {
	return s.encrypt(d)
}

func (s *session) Decrypt(d []byte) []byte {
	return s.decrypt(d)
}
