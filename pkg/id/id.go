package id

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
)

var k8sEncoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)

// GenerateID generates a random ID with the given prefix and length.
func GenerateID(prefix string, n int) (string, error) {
	return gen(prefix, n, base64.RawURLEncoding.EncodeToString)
}

// GenerateIDForK8SResource generates a random ID for a Kubernetes resource.
func GenerateIDForK8SResource(prefix string) (string, error) {
	return gen(prefix, 24, k8sEncoding.EncodeToString)
}

func gen(prefix string, n int, encodeFn func(src []byte) string) (string, error) {
	numBytes := (n * 3) / 4
	randBytes := make([]byte, numBytes)
	if _, err := rand.Read(randBytes); err != nil {
		return "", err
	}
	randStr := encodeFn(randBytes)
	if len(randStr) > n {
		randStr = randStr[:n]
	}
	return prefix + randStr, nil
}
