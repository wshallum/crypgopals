package set1

import (
	"encoding/hex"
)

// Given a byte string, try XOR-ing the byte string against every single byte
// value and return the value and the resulting buffer that most resembles English.
func GuessSingleByteXor(s string) (byte, []byte, error) {
	bin, err := hex.DecodeString(s)
	if err != nil {
		return 0, nil, err
	}

	// results: xor byte & score
	results := make([]byteIntPair, 256)
	buffer := make([]byte, len(bin))
	for b := 0; b < 256; b++ {
		for i := 0; i < len(buffer); i++ {
			buffer[i] = byte(b)
		}
		output, err := FixedXor(bin, buffer)
		if err != nil {
			return 0, nil, err
		}
		results[b].b = byte(b)
		results[b].i = ScoreString(output)
	}

	sortByteIntPairs(results, byteIntPairSortByIntDesc)

	for i := 0; i < len(buffer); i++ {
		buffer[i] = results[0].b
	}
	output, err := FixedXor(bin, buffer)
	if err != nil {
		return 0, nil, err
	}

	return results[0].b, output, nil
}
