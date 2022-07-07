package writer_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/matthieutran/leafre-login/networking/writer"
)

func TestWriteCheckUserLimitResult(t *testing.T) {
	var bWarningLevel, bPopulateLevel [1]byte
	rand.Read(bWarningLevel[:])
	rand.Read(bPopulateLevel[:])

	var b bytes.Buffer
	writer.WriteCheckUserLimitResult(&b, bWarningLevel[0], bPopulateLevel[0])

	// Ensure byte size == 4 (header + above packet)
	size := len(b.Bytes())
	if size != 4 {
		t.Error("Expected byte size == 4, received", size)
	}

	// Ensure header == OpcodeCheckUserLimitResult
	var header [2]byte
	_, err := b.Read(header[:])
	if err != nil {
		t.Fatal("Could not read header:", err)
	}
	if binary.LittleEndian.Uint16(header[:]) != writer.OpCodeCheckUserLimitResult {
		t.Errorf("Expected the packet header to be % X, got % X", writer.OpCodeCheckUserLimitResult, header)
	}

	// Ensure bWarningLevel is the same
	var actualBWarningLevel [1]byte
	_, err = b.Read(actualBWarningLevel[:])
	if err != nil {
		t.Fatal("Could not read bWarningLevel:", err)
	}
	if bWarningLevel != actualBWarningLevel {
		t.Errorf("Expected bWarningLevel to be % X, got % X", bPopulateLevel, actualBWarningLevel)
	}

	// Ensure bPopulateLevel is the same
	var actualBPopulateLevel [1]byte
	_, err = b.Read(actualBPopulateLevel[:])
	if err != nil {
		t.Fatal("Could not read actualBPopulateLevel:", err)
	}
	if bPopulateLevel != actualBPopulateLevel {
		t.Errorf("Expected bPopulateLevel to be % X, got % X", bPopulateLevel, actualBPopulateLevel)
	}
}
