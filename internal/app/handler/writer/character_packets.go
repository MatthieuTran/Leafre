package writer

import (
	"bytes"
	"io"

	"github.com/matthieutran/leafre-login/internal/domain/character"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

func WriteCharacterStats(w io.Writer, c character.Character) {
	pw := packet.NewPacketWriter()
	pw.WriteUInt32(c.ID)
	pw.WritePaddedString(c.Name, 13)

	pw.WriteOne(c.Gender)
	pw.WriteOne(c.Skin)
	pw.WriteUInt32(c.Face)
	pw.WriteUInt32(c.Hair)

	// TODO: Pets
	pw.WriteUInt32(0)
	pw.WriteUInt32(0)
	pw.WriteUInt32(0)

	pw.WriteOne(c.Level)
	pw.WriteUInt16(uint16(c.Job))
	pw.WriteUInt16(c.Strength)
	pw.WriteUInt16(c.Dexterity)
	pw.WriteUInt16(c.Intelligence)
	pw.WriteUInt16(c.Luck)
	pw.WriteUInt32(c.HP)
	pw.WriteUInt32(c.MaxHP)
	pw.WriteUInt32(c.MP)
	pw.WriteUInt32(c.MaxMP)

	pw.WriteUInt16(c.AP)
	if c.Job.IsExtendSPJob() {
		var spBuf bytes.Buffer
		WriteCharacterExtendSP(&spBuf, c)

		pw.WriteBytes(spBuf.Bytes())
	} else {
		pw.WriteUInt16(c.SP)
	}

	pw.WriteUInt32(c.Experience)
	pw.WriteUInt16(c.Popularity)

	pw.WriteUInt32(c.TempExperience)
	pw.WriteUInt32(c.FieldID)
	pw.WriteOne(c.FieldPortal)
	pw.WriteUInt32(c.PlayTime)
	pw.WriteUInt16(c.SubJob)

	w.Write(pw.Packet())
}

func WriteCharacterLook(w io.Writer, c character.Character) {
	pw := packet.NewPacketWriter()
	pw.WriteOne(c.Gender)
	pw.WriteOne(c.Skin)
	pw.WriteUInt32(c.Face)

	pw.WriteOne(0)
	pw.WriteUInt32(c.Hair)

	// Inventory
	for equip := range stickers {
		pw.WriteOne(key)
		pw.WriteUInt32(equip.ID)
	}
	pw.WriteOne(0xFF)
	for equip := range unseen {
		pw.WriteOne(key)
		pw.WriteUInt32(equip.ID)
	}
	pw.WriteOne(0xFF)
	// cash
	pw.WriteUInt32(0)

	for i := 0; i < 3; i++ {
		pw.WriteUInt32(0)
	}
}

func WriteCharacterExtendSP(w io.Writer, c character.Character) {
	pw := packet.NewPacketWriter()
	pw.WriteOne(byte(len(c.ExtendSP)))
	for jobLevel, sp := range c.ExtendSP {
		pw.WriteOne(jobLevel)
		pw.WriteOne(sp)
	}

	w.Write(pw.Packet())
}
