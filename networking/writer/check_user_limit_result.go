package writer

import (
	"io"

	"github.com/matthieutran/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeCheckUserLimitResult uint16 = 0x3

// WriteCheckUserLimitResult writes the world user limit information
func WriteCheckUserLimitResult(w io.Writer, bWarningLevel byte, bPopulateLevel byte) {
	p := packet.Packet{}
	p.WriteShort(OpCodeCheckUserLimitResult)
	p.WriteByte(bWarningLevel)
	p.WriteByte(bPopulateLevel)

	// Write world to client
	w.Write(p.Bytes())
}
