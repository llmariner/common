package id

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// GenerateID generates a random ID with the given prefix and length.
func GenerateID(prefix string, n int) (string, error) {
	numBytes := (n * 3) / 4
	randBytes := make([]byte, numBytes)
	if _, err := rand.Read(randBytes); err != nil {
		return "", err
	}

	randStr := base64.URLEncoding.EncodeToString(randBytes)
	randStr = strings.TrimRight(randStr, "=")

	if len(randStr) > n {
		randStr = randStr[:n]
	}
	return prefix + randStr, nil
}
