package announcer

import "github.com/matthieutran/packet"

func AnnounceHandshake(majorVersion uint16, minorVersion string, ivRecv, ivSend [4]byte, locale byte) []byte {
	p := packet.Packet{}
	p.WriteShort(14)            // Length of packet
	p.WriteShort(majorVersion)  // Maple Version (83)
	p.WriteString(minorVersion) // Subversion (1)
	p.WriteBytes(ivRecv[:])     // Recv IV
	p.WriteBytes(ivSend[:])     // Send IV
	p.WriteByte(locale)         // Maple Locale (8)

	return p.Bytes()
}
