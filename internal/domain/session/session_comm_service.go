package session

import (
	"context"
)

// SessionCommunicationService is a domain service that provides methods to communicate with the session
type SessionCommunicationService interface {
	EncryptPacket(ctx context.Context, id string, p []byte) (encrypted []byte, err error)
	DecryptPacket(ctx context.Context, id string, p []byte) (decrypted []byte, err error)
	ReadFromID(ctx context.Context, sessionID string, p []byte) (n int, err error)
	WriteToID(ctx context.Context, sessionID string, p []byte) (n int, err error)
}

func NewSessionCommunicationService(sr SessionRepository) SessionCommunicationService {
	return &defaultSessionCommunicationService{sessionRepo: sr}
}

type defaultSessionCommunicationService struct {
	sessionRepo SessionRepository
}

func (scs *defaultSessionCommunicationService) EncryptPacket(ctx context.Context, id string, p []byte) (encrypted []byte, err error) {
	s, err := scs.sessionRepo.GetById(ctx, id)
	if err != nil {
		return
	}

	return s.encrypt(p), nil
}

func (scs *defaultSessionCommunicationService) DecryptPacket(ctx context.Context, id string, p []byte) (decrypted []byte, err error) {
	s, err := scs.sessionRepo.GetById(ctx, id)
	if err != nil {
		return
	}

	return s.decrypt(p), nil
}

func (scs *defaultSessionCommunicationService) ReadFromID(ctx context.Context, sessionID string, p []byte) (n int, err error) {
	s, err := scs.sessionRepo.GetById(ctx, sessionID)
	if err != nil {
		return
	}

	return s.conn.Read(s.decrypt(p))
}

func (scs *defaultSessionCommunicationService) WriteToID(ctx context.Context, sessionID string, p []byte) (n int, err error) {
	s, err := scs.sessionRepo.GetById(ctx, sessionID)
	if err != nil {
		return
	}

	return s.conn.Write(s.encrypt(p))
}
