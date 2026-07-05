package main

import (
	"cryptopals/utils"
	"fmt"
)

func main() {
	h := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b := utils.DecodeHexToBytes(h)
	utf8 := utils.EncodeBytesToUtf8(b)
	b64 := utils.EncodeBytesToBase64(b)
	fmt.Printf("utf8=%s\nb64=%s\n", utf8, b64)
}
