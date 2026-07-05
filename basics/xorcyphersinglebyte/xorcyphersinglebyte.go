package xorcyphersinglebyte

import (
	"cryptopals/utils"
)

func score(b []byte) int {
	// https://www.commfront.com/pages/ascii-chart
	// 1. reasonable English characters only appear where first nibble in [2,7]
	// 2. we really like spaces
	s := 0
	for _, next_byte := range b {
		high_nib := next_byte >> 4
		if high_nib < 2 || high_nib > 7 {
			return -1
		}
		if high_nib >= 2 && high_nib <= 7 {
			s += 1
		}
		if next_byte == ' ' {
			s += 5
		}
	}
	return s
}

func make_candidate(b []byte, c byte) []byte {
	candidate := make([]byte, len(b))
	copy(candidate, b)
	for j := range candidate {
		candidate[j] ^= c
	}
	return candidate
}

func XORCypherSingleByte(h string) (byte, string) {
	b := utils.DecodeHexToBytes(h)
	highest_score := -1
	var best_byte byte
	var translated []byte
	for i := 0; i <= 255; i++ {
		candidate := make_candidate(b, byte(i))
		s := score(candidate)
		if s > highest_score {
			highest_score = s
			best_byte = byte(i)
			translated = candidate
		}
	}
	return best_byte, string(translated)
}
