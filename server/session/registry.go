package session

import (
	"errors"
	"net"
)

// SessionRegistry manages and keeps track of all the sessions
type SessionRegistry interface {
	Create(conn net.Conn, encrypter Codec, decrypter Codec) Session
	Add(Session) error
	Get(id string) (Session, error)
	Destroy(id string) error
}

func NewSessionRegistry() SessionRegistry {
	sessions := make(map[string]Session)

	return &sessionRegistry{sessions: sessions}
}

// sessionRegistry implements SessionRegistry with an in memory map
type sessionRegistry struct {
	sessions map[string]Session
}

var ErrAlreadyExists = errors.New("session already exists")
var ErrDoesNotExist = errors.New("session does not exist")

// Add adds a session to the registry
func (sr *sessionRegistry) Create(conn net.Conn, encrypter Codec, decrypter Codec) (s Session) {
	return NewSession(conn, encrypter, decrypter)
}

// Add adds a session to the registry
func (sr *sessionRegistry) Add(s Session) (err error) {
	if _, exists := sr.sessions[s.ID()]; exists {
		err = ErrAlreadyExists
		return
	}

	sr.sessions[s.ID()] = s

	return
}

func (sr *sessionRegistry) Get(id string) (s Session, err error) {
	s, exists := sr.sessions[id]
	if !exists {
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
