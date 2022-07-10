package reader_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/internal/app/handler/reader"
	"github.com/matthieutran/leafre-login/pkg/packet"
)

func randString(n int) string {
	var parts = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	res := make([]rune, n)
	for i := range res {
		res[i] = parts[rand.Intn(len(parts))]
	}

	return string(res)
}

func TestReadLoginMany(t *testing.T) {
	for i := 0; i < 100; i++ {
		TestReadLogin(t)
	}
}

func TestReadLogin(t *testing.T) {
	var user, pass string
	var machineId [16]byte
	var grc [4]byte
	var gsm [1]byte
	var unk1, unk2 [1]byte
	var pc [4]byte

	user = randString(rand.Intn(12))
	pass = randString(rand.Intn(12))
	rand.Read(machineId[:])
	rand.Read(grc[:])
	rand.Read(gsm[:])
	rand.Read(unk1[:])
	rand.Read(unk2[:])
	rand.Read(pc[:])

	pw := packet.NewPacketWriter()
	pw.WriteString(user)
	pw.WriteString(pass)
	pw.WriteBytes(machineId[:])
	pw.WriteBytes(grc[:])
	pw.WriteBytes(gsm[:])
	pw.WriteBytes(unk1[:])
	pw.WriteBytes(unk2[:])
	pw.WriteBytes(pc[:])

	res := reader.ReadLogin(pw.Packet())

	// Check if username is correct
	if res.Username != user {
		t.Errorf("Expected user == %s, actual user = %s", user, res.Username)
	}

	// Check if password is correct
	if res.Password != pass {
		t.Errorf("Expected pass == %s, actual pass = %s", pass, res.Password)
	}

	// Check if HWID is correct
	if !bytes.Equal(machineId[:], res.MachineId) {
		t.Errorf("Expected machineId == % X, actual machineId = % X", pass, res.Password)
	}

	// Check if GameRoomClient is correct
	numGrc := binary.LittleEndian.Uint32(grc[:])
	if numGrc != res.GameRoomClient {
		t.Errorf("Expected GameRoomClient == %d, actual GameRoomClient = %d", numGrc, res.GameRoomClient)
	}

	// Check if GameStartMode is correct
	if gsm[0] != res.GameStartMode {
		t.Errorf("Expected GameStartMode == %d, actual GameStartMode = %d", gsm[0], res.GameStartMode)
	}

	// Check if Unknown1 is correct
	if unk1[0] != res.Unknown1 {
		t.Errorf("Expected Unknown1 == %d, actual Unknown1 = %d", unk1[0], res.Unknown1)
	}

	// Check if Unknown2 is correct
	if unk2[0] != res.Unknown2 {
		t.Errorf("Expected Unknown2 == %d, actual Unknown2 = %d", unk2[0], res.Unknown2)
	}

	// Check if PartnerCode is correct
	pCode := binary.LittleEndian.Uint32(pc[:])
	if pCode != res.PartnerCode {
		t.Errorf("Expected PartnerCode == %d, actual PartnerCode = %d", pCode, res.PartnerCode)
	}
}
