package crypto

func encodeVersion(version uint16) (res uint16) {
	res = 0xFFFF - version
	res = (res >> 8 & 0xFF) | (res << 8 & 0xFF00)
	return
}
