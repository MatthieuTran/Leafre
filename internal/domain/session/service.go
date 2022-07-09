package session

import "context"

type SessionService interface {
	CreateSession(ctx context.Context, session Session) error
	RemoveSession(ctx context.Context, id string) error
	EncryptPacket(ctx context.Context, id string, p []byte) (encrypted []byte, err error)
	DecryptPacket(ctx context.Context, id string, p []byte) (decrypted []byte, err error)
	ReadFromID(ctx context.Context, sessionID string, p []byte) (n int, err error)
	WriteToID(ctx context.Context, sessionID string, p []byte) (n int, err error)
}

func NewSessionService(sr SessionRepository) SessionService {
	return &defaultSessionService{repository: sr}
}

type defaultSessionService struct {
	repository SessionRepository
}

func (ss *defaultSessionService) CreateSession(ctx context.Context, session Session) error {
	return ss.repository.Add(ctx, session)
}

func (ss *defaultSessionService) RemoveSession(ctx context.Context, id string) error {
	return ss.repository.Destroy(ctx, id)
}

func (ss *defaultSessionService) EncryptPacket(ctx context.Context, id string, p []byte) (encrypted []byte, err error) {
	s, err := ss.repository.GetById(ctx, id)
	if err != nil {
		return
	}

	return s.encrypt(p), nil
}

func (ss *defaultSessionService) DecryptPacket(ctx context.Context, id string, p []byte) (decrypted []byte, err error) {
	s, err := ss.repository.GetById(ctx, id)
	if err != nil {
		return
	}

	return s.decrypt(p), nil
}

func (ss *defaultSessionService) ReadFromID(ctx context.Context, sessionID string, p []byte) (n int, err error) {
	s, err := ss.repository.GetById(ctx, sessionID)
	if err != nil {
		return
	}

	return s.conn.Read(s.decrypt(p))
}

func (ss *defaultSessionService) WriteToID(ctx context.Context, sessionID string, p []byte) (n int, err error) {
	s, err := ss.repository.GetById(ctx, sessionID)
	if err != nil {
		return
	}

	return s.conn.Write(s.encrypt(p))
}
