package aws

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// KMSClient maintains the client to access AWS KMS related service.
type KMSClient struct {
	client *kms.Client
	// keyAlias the alias name of the KMS master key.
	keyAlias string
}

// NewKMSClient construct a KMS client.
func NewKMSClient(ctx context.Context, o NewConfigOptions, keyAlias string) (*KMSClient, error) {
	cfg, err := NewConfig(ctx, o)
	if err != nil {
		return nil, err
	}
	return &KMSClient{
		client:   kms.NewFromConfig(cfg),
		keyAlias: keyAlias,
	}, nil
}

// CreateDataKey creates a data key with KMS master key.
func (c *KMSClient) CreateDataKey(ctx context.Context) ([]byte, []byte, error) {
	keyOutput, err := c.client.GenerateDataKey(ctx, &kms.GenerateDataKeyInput{
		KeyId:   aws.String(c.keyAlias),
		KeySpec: types.DataKeySpecAes256,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("create data key: %s", err)
	}
	dataKey := keyOutput.Plaintext
	blob := keyOutput.CiphertextBlob
	return dataKey, blob, nil
}

// DecryptDataKey decrypts the data key with KMS master key.
func (c *KMSClient) DecryptDataKey(ctx context.Context, encryptedKey []byte) ([]byte, error) {
	decryptInput := &kms.DecryptInput{
		CiphertextBlob: encryptedKey,
	}

	decryptOutput, err := c.client.Decrypt(ctx, decryptInput)
	if err != nil {
		return nil, fmt.Errorf("decrypt data key: %s", err)
	}
	return decryptOutput.Plaintext, nil
}

// Encrypt encrypts a string data.
func Encrypt(ctx context.Context, data, name string, dataKey []byte) ([]byte, error) {
	// Create a cipher block and nounce.
	aesgcm, nounce, err := cipherHelper(dataKey, name)
	if err != nil {
		return nil, fmt.Errorf("encrypt: %s", err)
	}

	// Encrypt the data.
	blob := aesgcm.Seal(nil, nounce, []byte(data), nil)
	return blob, nil
}

// Decrypt decrypts into a string data.
func Decrypt(ctx context.Context, data []byte, name string, dataKey []byte) (string, error) {
	// Create a cipher block and nounce.
	aesgcm, nounce, err := cipherHelper(dataKey, name)
	if err != nil {
		return "", fmt.Errorf("decrypt: %s", err)
	}

	// Decrypt the data.
	plaintext, err := aesgcm.Open(nil, nounce, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("decrypt: %s", err)
	}
	return string(plaintext), nil
}

func cipherHelper(dataKey []byte, ID string) (cipher.AEAD, []byte, error) {
	// Create a cipher block using AES-GCM.
	block, err := aes.NewCipher(dataKey)
	if err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// Generate the nounce.
	// Since the open and seal must use the same nounce, generating nounce based on user ID.
	len := aesgcm.NonceSize()
	nounce := []byte(ID)[:len]
	return aesgcm, nounce, nil
}
