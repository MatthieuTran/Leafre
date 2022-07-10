package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

// WorldInformation provides information about each available world to the client
var OpCodeCheckUserLimitResult uint16 = 0x3

type SendCheckUserLimit struct {
	WarningLevel  byte
	PopulateLevel byte
}

// WriteCheckUserLimitResult writes the world user limit information
func WriteCheckUserLimitResult(w io.Writer, send SendCheckUserLimit) {
	pw := packet.NewPacketWriter()
	pw.WriteUInt16(OpCodeCheckUserLimitResult)
	pw.WriteOne(send.WarningLevel)
	pw.WriteOne(send.PopulateLevel)

	// Write world to client
	w.Write(pw.Packet())
}
