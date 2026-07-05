package main

import (
	"cryptopals/basics/xorcyphersinglebyte"
	"fmt"
)

func main() {
	h := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	single, ans := xorcyphersinglebyte.XORCypherSingleByte(h)
	fmt.Printf("byte_hex=0x%02x (%s)\n", single, ans)
}
