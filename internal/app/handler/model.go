package handler

import (
	"fmt"
	"io"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

var OpCodeEXAMPLE = 0x0

// A PacketHandler will provide a Handle method that takes in a session, eventstreamer, and a byte slice. It should also implement the Stringer interface.
//
// While not in the interface, the PacketHandler should also provide an Opcode to identify itself in the same file.
type PacketHandler interface {
	Handle(io.Writer, packet.Packet)
	fmt.Stringer
}
