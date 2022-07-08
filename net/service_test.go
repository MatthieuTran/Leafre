package net_test

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"testing"

	netLogin "github.com/matthieutran/leafre-login/net"
)

// TestHandleConn tests the HandleConn function
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
func TestConnHandshake(t *testing.T) {
	majorVersion := 95
	minorVersion := "1"
	locale := 8

	conn1, conn2 := net.Pipe()
	fmt.Println("TEST")
	go netLogin.HandleConn(conn1)

	// Client should receive a handshake packet
	buf := make([]byte, 1024)
	conn2.Read(buf)

	// Packet Size == 00 14
	partSize := []byte{14, 0}
	if !bytes.Equal(buf[0:2], partSize) {
		t.Errorf("Expected buf[0:2] (packet size) == %d, actual = %d", partSize, buf[0:2])
	}

	// majorVersion == 95 00
	partVersion := []byte{byte(majorVersion), 0}
	if !bytes.Equal(buf[2:4], partVersion) {
		t.Errorf("Expected buf[2:4] (majorVersion) == %d, actual = %d", partVersion, buf[2:4])
	}

	// minorVersion == 1 0 49
	partMinorVersion := []byte{byte(len(minorVersion)), 0}
	partMinorVersion = append(partMinorVersion, []byte(minorVersion)...)
	if !bytes.Equal(buf[4:7], partMinorVersion) {
		t.Errorf("Expected handshake packet minorVersion == %d, actual = %d", partMinorVersion, buf[4:7])
	}

	// Skip the next 8 bytes since recvIV and sendIV are random
	// locale == 8
	if buf[15] != byte(locale) {
		t.Errorf("Expected locale == %d, actual = %d", locale, buf[15])
	}

	log.Println("END")

	conn1.Close()
	conn2.Close()
}
