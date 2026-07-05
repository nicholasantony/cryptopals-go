package fixedxor

import (
	"crypto/subtle"
	"cryptopals/utils"
)

func XORTwoHexBuffers(s1 string, s2 string) string {
	b1 := utils.DecodeHexToBytes(s1)
	b2 := utils.DecodeHexToBytes(s2)
	xor := make([]byte, len(b1))
	subtle.XORBytes(xor, b1, b2)
	return utils.EncodeBytesToHex(xor)
}
