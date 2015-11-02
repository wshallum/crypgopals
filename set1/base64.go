package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(s string) (string, error) {
	bin, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bin), nil
}
