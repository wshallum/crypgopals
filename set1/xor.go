package set1

import "errors"

var (
	ErrBufferLengthMismatch = errors.New("FixedXor: buffers not of same length")
	ErrBufferNil            = errors.New("FixedXor: a buffer is nil")
)

func FixedXor(a, b []byte) ([]byte, error) {
	if a == nil || b == nil {
		return nil, ErrBufferNil
	}
	if len(a) != len(b) {
		return nil, ErrBufferLengthMismatch
	}
	result := make([]byte, len(a))
	for i, aByte := range a {
		bByte := b[i]
		result[i] = aByte ^ bByte
	}
	return result, nil
}
