package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func NewAesGCM(key []byte) *aesGCM {
	return &aesGCM{
		key:          key,
		stdNonceSize: 12,
	}
}

// AES-256 GCM Utils
type aesGCM struct {
	key          []byte
	stdNonceSize int
}

func (c *aesGCM) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, c.stdNonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	gcmCipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := gcmCipher.Seal(nil, nonce, plainText, nil)
	return append(nonce, cipherText...), nil
}

func (c *aesGCM) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	gcmCipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := cipherText[0:c.stdNonceSize]
	plainText, err := gcmCipher.Open(nil, nonce, cipherText[c.stdNonceSize:], nil)
	if err != nil {
		return nil, err
	}

	return plainText, err
}

func (c *aesGCM) EncryptString(plainText string) (string, error) {
	encrypted, err := c.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (c *aesGCM) DecryptString(cipherText string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	decrypted, err := c.Decrypt(b)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
