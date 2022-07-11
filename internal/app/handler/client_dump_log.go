package handler

import (
	"log"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

var OpCodeClientDumpLog uint16 = 0x24

type HandlerClientDumpLog struct{}

func NewHandlerClientDumpLog() HandlerClientDumpLog {
	return HandlerClientDumpLog{}
}

func (h *HandlerClientDumpLog) Handle(s session.Session, p packet.Packet) {
	recv := reader.ReadClientDumpLog(p)
	log.Printf("Client exited with error code: %d (call type: %s) from operation 0x%X with payload: %s (SeqSend: %d)", recv.ErrorCode, recv.Type, recv.Operation, packet.Packet(recv.Payload), recv.SeqSend)
}

func (h *HandlerClientDumpLog) String() string {
	return "ClientDumpLog"
}
