package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func EncodeBytesToUtf8(b []byte) string {
	return string(b)
}

func EncodeBytesToBase64(b []byte) string {
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(b64, b)
	return string(b64)
}

func EncodeBytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

func DecodeHexToBytes(h string) []byte {
	// converts h, a hex-encoded string, to raw bytes
	b, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return b
}

func DecodeBase64ToBytes(b64 string) []byte {
	// converts b64, a base64-encoded string, to raw bytes
	b, err := base64.RawStdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}
	return b
}

func DecodeUft8ToBytes(h string) []byte {
	return []byte(h)
}
