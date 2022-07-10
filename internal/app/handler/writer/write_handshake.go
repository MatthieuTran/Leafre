package writer

import (
	"io"

	"github.com/matthieutran/leafre-login/pkg/packet"
)

func WriteHandshake(w io.Writer) func(majorVersion uint16, minorVersion string, ivRecv, ivSend []byte, locale byte) {
	return func(majorVersion uint16, minorVersion string, ivRecv, ivSend []byte, locale byte) {
		pw := packet.NewPacketWriter()
		pw.WriteUInt16(14)           // Packet Length
		pw.WriteUInt16(majorVersion) // Maple Version
		pw.WriteString(minorVersion) // Subversion
		pw.WriteBytes(ivRecv)        // ivRecv
		pw.WriteBytes(ivSend)        // ivSend
		pw.WriteOne(locale)          // Locale (8)

		w.Write(pw.Packet())
	}
}
