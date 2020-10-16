package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncAesGcm returns data encrypted with AES
// The key argument should be the AES key,
// either 16 or 32 bytes to select
// AES-128 or AES-256.
func EncAesGcm(data, key, nonce []byte) ([]byte, error) {
	k := len(key)
	switch k {
	case 16, 32:
		break
	default:
		return nil, aes.KeySizeError(k)
	}
	return gcmEncrypt(data, key, nonce)
}

func DecAesGcm(data, key, nonce []byte) ([]byte, error) {
	k := len(key)
	switch k {
	case 16, 32:
		break
	default:
		return nil, aes.KeySizeError(k)
	}

	return gcmDecrypt(data, key, nonce)
}

func gcmEncrypt(data, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesGcm.Seal(nil, nonce, data, nil)
	return ciphertext, nil
}

func gcmDecrypt(encData, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGcm.Open(nil, nonce, encData, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
