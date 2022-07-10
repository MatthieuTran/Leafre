package inmem_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/matthieutran/leafre-login/internal/adapters/inmem"
	"github.com/matthieutran/leafre-login/internal/domain/session"
)

func TestSessionRepository(t *testing.T) {
	ctx := context.Background()
	conn, _ := net.Pipe()
	mockCodec := func(d []byte) []byte { return d }
	r := inmem.NewSessionRepository()
	s := session.NewSession(conn, mockCodec, mockCodec)

	ctxAdd, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err := r.Add(ctxAdd, s)
	if err != nil {
		t.Error("Cannot add session to repository:", err)
	}

	ctxGet, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	_, err = r.GetByID(ctxGet, s.ID())
	if err != nil {
		t.Error("Cannot get session from repository:", err)
	}

	ctxDestroy, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err = r.Destroy(ctxDestroy, s.ID())
	if err != nil {
		t.Error("Cannot destroy session from repository:", err)
	}
}
