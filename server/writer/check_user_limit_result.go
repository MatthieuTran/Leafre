package writer

import (
	"io"

	"github.com/matthieutran/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeCheckUserLimitResult uint16 = 0x3

// WriteCheckUserLimitResult writes the world user limit information
func WriteCheckUserLimitResult(w io.Writer) {
	p := packet.Packet{}
	p.WriteShort(OpCodeCheckUserLimitResult)
	p.WriteByte(0) // bWarningLevel
	p.WriteByte(0) // bPopulateLevel

	// Write world to client
	w.Write(p.Bytes())
}
