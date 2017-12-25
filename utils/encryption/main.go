package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

func srringToEncryption(s string) string {
	converted := sha256.Sum256([]byte(s))

	return hex.EncodeToString(converted[:])
}

// GetPassword get password
func GetPassword(p string, k string) string {
	return srringToEncryption(p + k)
}
