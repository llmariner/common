package aws

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptAndDecrypt(t *testing.T) {
	ctx := context.TODO()
	dataKey := []byte{189, 210, 91, 253, 139, 50, 99, 24, 216, 87, 255, 114, 240, 71, 14, 182, 104, 251, 46, 62, 206, 22, 229, 117, 121, 102, 115, 81, 72, 99, 43, 194}
	s := fmt.Sprintf("%d:%s:%s", 1, "client Id123", "my client secret")
	apiKey := base64.StdEncoding.EncodeToString([]byte(s))
	tid := "tenant-uuid1"

	encrypted, err := Encrypt(ctx, apiKey, tid, dataKey)
	assert.NoError(t, err)

	decrypted, err := Decrypt(ctx, encrypted, tid, dataKey)
	assert.NoError(t, err)
	assert.Equal(t, apiKey, decrypted)
}
