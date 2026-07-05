package hex2base64_test

import (
	"cryptopals/basics/hex2base64"
	"testing"
)

func TestHex2Base64(t *testing.T) {
	hex_in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	t.Logf("running hex to base64 test: %s\n", hex_in)
	utf8_english := "I'm killing your brain like a poisonous mushroom"
	t.Logf("in english: %s\n", utf8_english)
	ans_b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if ret := hex2base64.ConvertHexToBase64(hex_in); ret != ans_b64 {
		t.Logf("expected result: %s\n", ans_b64)
		t.Logf("actual result: %s\n", ret)
		t.FailNow()
	}
	t.Logf("test case pass!")
}
