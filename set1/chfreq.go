package set1

import (
	"strings"
)

// This returns a map[byte]uint that maps bytes to
// how many times the byte occurs in string s.
// This operates on BYTES, not runes.
func ByteFrequencies(s []byte) map[byte]int {
	freqs := make(map[byte]int, 256)
	for i := 0; i < 256; i++ {
		freqs[byte(i)] = 0
	}
	for i := 0; i < len(s); i++ {
		freqs[s[i]] += 1
	}
	return freqs
}

// ScoreString scores a string's resemblance to English on
// the basis of byte frequencies. The higher the score,
// the closer it looks to English.
func ScoreString(s []byte) int {
	// how to score: we sort the letter frequencies then compare against the english letter frequencies
	// we then compute sum of difference between the rank of the letter in english and the rank of the
	// letter in text, then negate it so the closer the ranks are, the higher the score
	//
	// We then subtract one point for every non letter character
	const order = "etaoinsrhdlucmfywgpbvkxqjz"
	freqs := ByteFrequencies(s)
	letterFreqs := make(map[byte]byteIntPair, 26)
	for i := 0; i < len(order); i++ {
		b := order[i]
		letterFreqs[b] = byteIntPair{b: b, i: freqs[b] + freqs[b-byte(32)]}
	}
	nonLetterTotal := 0
	for b, f := range freqs {
		if (b < 32 || b > 127) && f > 0 {
			nonLetterTotal += 100
		}
		if !((b >= byte('a') && b <= byte('z')) || (b >= byte('A') && b <= byte('Z'))) {
			nonLetterTotal += f
		}
	}
	orderedFreqs := make([]byteIntPair, 0)
	for _, f := range letterFreqs {
		if f.i > 0 {
			orderedFreqs = append(orderedFreqs, f)
		}
	}
	sortByteIntPairs(orderedFreqs, byteIntPairSortByIntDesc)
	result := 0
	for i, f := range orderedFreqs {
		dist := i - strings.IndexByte(order, f.b)
		if dist < 0 {
			dist = -dist
		}
		result -= dist
	}
	result -= nonLetterTotal
	return result
}
