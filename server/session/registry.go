package session

import (
	"errors"
	"net"
)

type SessionRegistry interface {
	Create(conn net.Conn, encrypter Codec, decrypter Codec) (Session, error)
	Get(id string) (Session, error)
	Destroy(id string) error
}

func NewSessionRegistry() SessionRegistry {
	sessions := make(map[string]Session)

	return &sessionRegistry{sessions: sessions}
}

type sessionRegistry struct {
	sessions map[string]Session
}

var ErrAlreadyExists = errors.New("session already exists")
var ErrDoesNotExist = errors.New("session does not exist")

func (sr *sessionRegistry) formID(conn net.Conn) string {
	return conn.RemoteAddr().String()
}

func (sr *sessionRegistry) Create(conn net.Conn, encrypter, decrypter Codec) (s Session, err error) {
	if _, exists := sr.sessions[sr.formID(conn)]; exists {
		err = ErrAlreadyExists
		return
	}

	sr.sessions[sr.formID(conn)] = newSession(conn, encrypter, decrypter)
	s = sr.sessions[sr.formID(conn)]

	return
}

func (sr *sessionRegistry) Get(id string) (s Session, err error) {
	if _, exists := sr.sessions[id]; !exists {
		err = ErrDoesNotExist
		return
	}
	return
}

func (sr *sessionRegistry) Destroy(id string) (err error) {
	if _, exists := sr.sessions[id]; !exists {
		err = ErrDoesNotExist
		return
	}

	delete(sr.sessions, id)
	return
}
