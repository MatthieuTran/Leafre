package session

import (
	"net"

	"github.com/matthieutran/leafre-login/internal/domain/user"
)

type Codec func(d []byte) []byte

type Session struct {
	id      string
	conn    net.Conn
	encrypt Codec
	decrypt Codec
	Account user.User
}

func NewSession(conn net.Conn, encrypter, decrypter Codec) Session {
	return Session{id: conn.RemoteAddr().String(), conn: conn, encrypt: encrypter, decrypt: decrypter}
}

func (s Session) ID() string {
	return s.conn.RemoteAddr().String()
}

func (s Session) Write(p []byte) (n int, err error) {
	return s.conn.Write(s.encrypt(p))
}

func (s Session) Read(p []byte) (n int, err error) {
	return s.conn.Read(s.decrypt(p))
}

func (s Session) Encrypt(p []byte) []byte {
	return s.encrypt(p)
}

func (s Session) Decrypt(p []byte) []byte {
	return s.decrypt(p)
}

func (s Session) SetAccount(u user.User) Session {
	var res Session
	res.Account = u
	return res
}
