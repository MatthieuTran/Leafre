package session

import (
	"net"
)

type Codec func(d []byte) []byte

type Session struct {
	id      string
	conn    net.Conn
	encrypt Codec
	decrypt Codec
}

func NewSession(conn net.Conn, encrypter, decrypter Codec) Session {
	return Session{id: conn.RemoteAddr().String(), conn: conn, encrypt: encrypter, decrypt: decrypter}
}

func (s Session) ID() string {
	return s.conn.RemoteAddr().String()
}
