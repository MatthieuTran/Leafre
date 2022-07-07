package reader_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/networking/reader"
	"github.com/matthieutran/packet"
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

	p := packet.Packet{}
	p.WriteString(user)
	p.WriteString(pass)
	p.WriteBytes(machineId[:])
	p.WriteBytes(grc[:])
	p.WriteBytes(gsm[:])
	p.WriteBytes(unk1[:])
	p.WriteBytes(unk2[:])
	p.WriteBytes(pc[:])

	res := reader.ReadLogin(p)

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
	numGrc := int(binary.LittleEndian.Uint32(grc[:]))
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
	pCode := int(binary.LittleEndian.Uint32(pc[:]))
	if pCode != res.PartnerCode {
		t.Errorf("Expected PartnerCode == %d, actual PartnerCode = %d", pCode, res.PartnerCode)
	}
}
