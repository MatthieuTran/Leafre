package handler

import (
	"github.com/matthieutran/duey"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

type PacketHandler interface {
	Name() string
	Handle(*tcpserve.Session, *duey.EventStreamer, packet.Packet) []byte
}
