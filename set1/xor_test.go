package set1

import (
	"encoding/hex"
	"testing"
)

func TestFixedXor(t *testing.T) {
	cases := []struct {
		a, b, result string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}

	for _, c := range cases {
		a, err := hex.DecodeString(c.a)
		if err != nil {
			t.Errorf("hex.DecodeString(%q) returned error", c.a)
		}
		b, err := hex.DecodeString(c.b)
		if err != nil {
			t.Errorf("hex.DecodeString(%q) returned error", c.b)
		}
		result, err := FixedXor(a, b)
		if err != nil {
			t.Errorf("FixedXor(%q, %q) returned error", c.a, c.b)
		}
		hexResult := hex.EncodeToString(result)
		if hexResult != c.result {
			t.Errorf("FixedXor(%q, %q) == %q, want %q", c.a, c.b, hexResult, c.result)
		}
	}

}
