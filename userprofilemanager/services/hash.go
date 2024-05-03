package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
