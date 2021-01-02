package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func NewAesCBC(key []byte) *aesCBC {
	return &aesCBC{
		key: key,
	}
}

// AES-256 GCM Utils
type aesCBC struct {
	key []byte
}

func (c *aesCBC) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}
	plainText = pkcsPadding(plainText, block.BlockSize())
	cipherText := make([]byte, block.BlockSize()+len(plainText))

	// Generate IV
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Crypt blocks blocks
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherText[block.BlockSize():], plainText)
	return cipherText, nil
}

func (c *aesCBC) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < block.BlockSize() {
		return nil, errors.New("data is too short")
	}

	iv := cipherText[:block.BlockSize()]
	cipherText = cipherText[block.BlockSize():]
	if len(cipherText)%block.BlockSize() != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return pkcsStripPadding(cipherText), nil
}

func (c *aesCBC) EncryptString(plainText string) (string, error) {
	encrypted, err := c.Encrypt([]byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (c *aesCBC) DecryptString(cipherText string) (string, error) {
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
