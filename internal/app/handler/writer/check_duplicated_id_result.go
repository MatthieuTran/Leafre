package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

type SendDuplicatedIDResult struct {
	Name      string
	Duplicate byte
}

var OpCodeCheckDuplicatedIDResult uint16 = 0xD

func WriteCheckDuplicatedIDResult(w io.Writer, send SendDuplicatedIDResult) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(OpCodeCheckDuplicatedIDResult)
	p.WriteString(send.Name)
	p.WriteOne(send.Duplicate)

	w.Write(p.Packet())
}
