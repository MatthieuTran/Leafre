package reader

import "github.com/matthieutran/leafre-login/pkg/packet"

type RecvClientDumpLog struct {
	Type      CallType
	ErrorCode uint32
	SeqSend   uint32
	Operation uint16
	Payload   []byte
}

type CallType uint16

func (t CallType) String() string {
	switch t {
	case 1: // CInPacket::Decode
		return "CInPacket::Decode"
	case 2: // CTLException
		return "CTLException"
	case 3: // CMSException
		return "CMSException"
	default:
		return "Unknown"
	}
}

func ReadClientDumpLog(p packet.Packet) (recv RecvClientDumpLog) {
	pr := packet.NewPacketReader(p)
	recv.Type = CallType(pr.ReadUInt16())
	recv.ErrorCode = pr.ReadUInt32()
	backupBufferSize := pr.ReadUInt16()
	backupBuffer := pr.ReadBytes(int(backupBufferSize))

	backupPacket := packet.NewPacketReader(backupBuffer)
	recv.SeqSend = backupPacket.ReadUInt32()
	recv.Operation = backupPacket.ReadUInt16()
	recv.Payload = backupPacket.ReadAvailableBytes()

	return
}
