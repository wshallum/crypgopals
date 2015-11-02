package set1

import "testing"

func TestHexToBase64(t *testing.T) {
	cases := []struct {
		hex, base64 string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}

	for _, c := range cases {
		result, err := HexToBase64(c.hex)
		if err != nil {
			t.Errorf("HexToBase64(%q) returned error", c.hex)
		}
		if result != c.base64 {
			t.Errorf("HexToBase64(%q) == %q, want %q", c.hex, result, c.base64)
		}
	}
}
