package set1

import "testing"

func TestByteFrequencies(t *testing.T) {
	f := ByteFrequencies([]byte("abcbcc"))
	if f[byte('a')] != 1 || f[byte('b')] != 2 || f[byte('c')] != 3 || f[byte('d')] != 0 {
		t.Errorf("Unexpected")
	}
}

func TestScoreString(t *testing.T) {
	var score int
	score = ScoreString([]byte("eeeetttaao"))
	if score != 0 {
		t.Errorf("Unexpected: %d", score)
	}
}
