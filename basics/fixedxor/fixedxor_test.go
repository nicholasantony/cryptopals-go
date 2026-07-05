package fixedxor_test

import (
	"cryptopals/basics/fixedxor"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	t.Logf("running test xor two hex strings:\n")
	s1 := "1c0111001f010100061a024b53535009181c"
	t.Logf("s1: %s\n", s1)
	s2 := "686974207468652062756c6c277320657965"
	t.Logf("s2: %s\n", s2)
	ans := "746865206b696420646f6e277420706c6179"
	if ret := fixedxor.XORTwoHexBuffers(s1, s2); ret != ans {
		t.Logf("expected result: %s\n", ans)
		t.Logf("actual result: %s\n", ret)
		t.FailNow()
	}
	t.Logf("test case pass!")
}
