package writer

import (
	"io"
	"log"

	"github.com/matthieutran/leafre-login/net/packet"
)

func WriteHandshake(w io.Writer, majorVersion uint16, minorVersion string, ivRecv, ivSend []byte, locale byte) {
	p := packet.NewPacketWriter()
	p.WriteUInt16(14)           // Packet Length
	p.WriteUInt16(majorVersion) // Maple Version
	p.WriteString(minorVersion) // Subversion
	p.Write(ivRecv)             // ivRecv
	p.Write(ivSend)             // ivSend
	p.WriteByte(locale)         // Locale (8)
	log.Println(p.Packet())

	w.Write(p.Packet())
}
