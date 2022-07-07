package writer_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/matthieutran/leafre-login/networking/writer"
	"github.com/matthieutran/leafre-login/user"
)

// TestWriteCheckPasswordResultNotSuccess tests that the writer should receive the auth fail result only.
func TestWriteCheckPasswordResultNotSuccess(t *testing.T) {
	var b bytes.Buffer

	writer.WriteCheckPasswordResult(&b, user.LoginResponseAuthFail, user.User{})

	// Ensure size <= 8
	size := len(b.Bytes())
	if size > 8 {
		t.Error("Expected packet of size <= 8, got", size)
	}

	// Ensure header = OpcodeCheckPassword
	var header [2]byte
	_, err := b.Read(header[:])
	if err != nil {
		t.Fatal("Could not read header:", err)
	}
	if binary.LittleEndian.Uint16(header[:]) != writer.OpCodeCheckPasswordResult {
		t.Errorf("Expected the packet header to be % X, got % X", writer.OpCodeCheckPasswordResult, header)
	}

	// Ensure code = AuthFail
	code, err := b.ReadByte()
	if err != nil {
		t.Fatal("Could not read bytes:", err)
	}
	if code != byte(user.LoginResponseAuthFail) {
		t.Errorf("Expected the result code to be % X, got % X", user.LoginResponseAuthFail, code)
	}
}

// TestWriteCheckPasswordResultSuccess tests that the writer should receive a success result with the user details
func TestWriteCheckPasswordResultSuccess(t *testing.T) {
	var b bytes.Buffer
	id := 500

	writer.WriteCheckPasswordResult(&b, user.LoginResponseSuccess, user.User{Id: id})

	// Ensure packet size > 8
	fmt.Printf("% X\n", b)
	size := len(b.Bytes())
	if size < 9 {
		t.Error("Expected packet of size > 8, got", size)
	}

	// Ensure header = OpCodeCheckPassword
	var header [2]byte
	_, err := b.Read(header[:])
	if err != nil {
		t.Fatal("Could not read header:", err)
	}
	if binary.LittleEndian.Uint16(header[:]) != writer.OpCodeCheckPasswordResult {
		t.Errorf("Expected the packet header to be % X, got % X", writer.OpCodeCheckPasswordResult, header)
	}

	// Ensure code = Success
	code, err := b.ReadByte()
	if err != nil {
		t.Fatal("Could not read bytes:", err)
	}
	if code != byte(user.LoginResponseSuccess) {
		t.Errorf("Expected the result code to be % X, got % X", user.LoginResponseAuthFail, code)
	}

	// Skip bytes
	var dummy [5]byte
	_, err = b.Read(dummy[:])
	if err != nil {
		t.Fatal("Could not read bytes:", err)
	}

	// Ensure ID = 500
	var raw_id [4]byte
	_, err = b.Read(raw_id[:])
	if err != nil {
		t.Fatal("Could not read bytes:", err)
	}
	pId := binary.LittleEndian.Uint32(raw_id[:])
	if pId != uint32(id) {
		t.Errorf("Expected ID to be %d, got %d", id, pId)
	}
}
