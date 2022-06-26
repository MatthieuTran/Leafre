package codec

import (
	"github.com/matthieutran/crypto"
)

func GenerateCodecs(version int, ivRecv, ivSend [4]byte) (encrypter, decrypter func(d []byte)) {

	// Create codecs
	recv := crypto.NewCodec(ivRecv, version)
	send := crypto.NewCodec(ivSend, version)

	// Create encrypter
	encrypter = func(d []byte) {
		send.Encrypt(d, true, true)
	}

	// Create decrypter
	decrypter = func(d []byte) {
		recv.Decrypt(d, true, true)
	}

	return
}
