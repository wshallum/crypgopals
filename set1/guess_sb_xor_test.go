package set1

import "testing"

func TestGuessSingleByteXor(t *testing.T) {
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	b, out, err := GuessSingleByteXor(s)
	if err != nil {
		t.Errorf("Failed %s", err.Error())
	}
	if b != 88 {
		t.Errorf("Failed %d %s", b, string(out))
	}

}
