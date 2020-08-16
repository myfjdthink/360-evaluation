package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str []byte) string {
	h := sha256.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
