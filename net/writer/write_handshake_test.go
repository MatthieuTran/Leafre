package writer_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/net/writer"
)

// TestWriteHandshake tests the writing part of the handshake to ensure the structure checks out
//
// Handshake packet details:
// Example handshake: [14 00] [83 00] [1] [00 49] [0 0 0 0] [0 0 0 0] [8]
// [14 00]   - Size of packet proceeding
// [83 00]   - majorVersion (the "Maple" version)
// [1]       - Size of the upcoming string
// [00 49]   - minorVersion (subversion is represented as a string)
// [0 0 0 0] - recvIV (this number is randomly generated and used for encryption/decryption)
// [0 0 0 0] - sendIV (this number is randomly generated and used for encryption/decryption)
// [8]       - locale (GMS = 8)
func TestWriteHandshake(t *testing.T) {
	var b bytes.Buffer

	var majorVersion [2]byte
	var minorVersion string
	var minorVersionSize int
	var ivRecv [4]byte
	var ivSend [4]byte
	var locale [1]byte

	rand.Read(majorVersion[:])
	minorVersionSize = rand.Intn(16)
	minorVersion = randString(minorVersionSize)
	rand.Read(ivRecv[:])
	rand.Read(ivSend[:])
	rand.Read(locale[:])

	writer.WriteHandshake(&b, binary.LittleEndian.Uint16(majorVersion[:]), minorVersion, ivRecv[:], ivSend[:], locale[0])
	res := b.Bytes()

	indx := 2

	// Check majorVersion
	if !bytes.Equal(majorVersion[:], res[indx:indx+2]) {
		t.Errorf("Expected majorVersion == %d, actual = %d", majorVersion, res[indx:indx+2])
	}
	indx += 2

	// Check minorVersion Size
	actualMinorVersionSize := res[indx]
	indx += 2
	if actualMinorVersionSize != byte(minorVersionSize) {
		t.Errorf("Expected minorVersionSize == %d, actual = %d", actualMinorVersionSize, minorVersionSize)
	}
	// Check minorVersion
	actualMinorVersion := res[indx : indx+int(actualMinorVersionSize)]
	indx += int(actualMinorVersionSize)
	if !bytes.Equal(actualMinorVersion, []byte(minorVersion)) {
		t.Errorf("Expected minorVersion == %d, actual = %d", actualMinorVersion, []byte(minorVersion))
	}

	// Check ivRecv
	actualIVRecv := res[indx : indx+4]
	indx += 4
	if !bytes.Equal(actualIVRecv, ivRecv[:]) {
		t.Errorf("Expected ivRecv == %d, actual = %d", actualIVRecv, ivRecv)
	}

	// Check ivSend
	actualIVSend := res[indx : indx+4]
	indx += 4
	if !bytes.Equal(actualIVSend, ivSend[:]) {
		t.Errorf("Expected ivSend == %d, actual = %d", actualIVSend, ivSend)
	}

	// Check locale
	actualLocale := res[indx]
	indx++
	if actualLocale != locale[0] {
		t.Errorf("Expected locale == %d, actual = %d", actualLocale, locale)
	}
}

func randString(n int) string {
	var parts = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	res := make([]rune, n)
	for i := range res {
		res[i] = parts[rand.Intn(len(parts))]
	}

	return string(res)
}
