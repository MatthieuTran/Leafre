package handler

import (
	"fmt"

	"github.com/matthieutran/leafre-login/internal/domain/session"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

var OpCodeEXAMPLE = 0x0

// A PacketHandler will provide a Handle method that takes in a session, eventstreamer, and a byte slice. It should also implement the Stringer interface.
//
// While not in the interface, the PacketHandler should also provide an Opcode to identify itself in the same file.
type PacketHandler interface {
	Handle(session.Session, packet.Packet)
	fmt.Stringer
}
