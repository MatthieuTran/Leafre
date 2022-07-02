package handler

import (
	"fmt"

	"github.com/matthieutran/duey"
	"github.com/matthieutran/packet"
	"github.com/matthieutran/tcpserve"
)

// A PacketHandler will provide a Handle method that takes in a session, eventstreamer, and a byte slice. It should also implement the Stringer interface.
//
// While not in the interface, the PacketHandler should also provide an Opcode to identify itself in the same file.
type PacketHandler interface {
	Handle(*tcpserve.Session, *duey.EventStreamer, packet.Packet)
	fmt.Stringer
}
