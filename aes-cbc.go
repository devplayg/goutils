package goutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncAesCbc returns data encrypted with AES
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func EncAesCbc(data, key []byte) ([]byte, error) {
	k := len(key)
	switch k {
	case 16, 24, 32:
		break
	default:
		return nil, aes.KeySizeError(k)
	}
	return cbcEncrypt(data, key)
}

func DecAesCbc(data, key []byte) ([]byte, error) {
	k := len(key)
	switch k {
	case 16, 24, 32:
		break
	default:
		return nil, aes.KeySizeError(k)
	}

	return cbcDecrypt(data, key)
}

// EncAes256Cbc returns data encrypted with AES-128
func EncAes256Cbc(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 32)
}

// EncAes192Cbc returns data encrypted with AES-192
func EncAes192Cbc(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 24)
}

// EncAes128Cbc returns data encrypted with AES-256
func EncAes128Cbc(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 16)
}

// DecAes256Cbc returns data decrypted with AES-128
func DecAes256Cbc(data, key []byte) ([]byte, error) {
	return decAESWithKeyLen(data, key, 32)
}

// DecAes192Cbc returns data decrypted with AES-192
func DecAes192Cbc(data, key []byte) ([]byte, error) {
	return decAESWithKeyLen(data, key, 24)
}

// DecAes128Cbc returns data decrypted with AES-256
func DecAes128Cbc(data, key []byte) ([]byte, error) {
	return decAESWithKeyLen(data, key, 16)
}

func encAESWithKeyLen(data, key []byte, keyLen int) ([]byte, error) {
	if len(key) != keyLen {
		return nil, aes.KeySizeError(keyLen)
	}
	return cbcEncrypt(data, key)
}

func decAESWithKeyLen(data, key []byte, keyLen int) ([]byte, error) {
	if len(key) != keyLen {
		return nil, aes.KeySizeError(keyLen)
	}
	return cbcDecrypt(data, key)
}

func pkcsPadding(data []byte, blockSize int) []byte {
	paddingLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(data, padding...)
}

func pkcsStripPadding(data []byte) []byte {
	length := len(data)
	paddingLen := int(data[length-1])
	return data[:(length - paddingLen)]
}

func cbcEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = pkcsPadding(data, block.BlockSize())
	encData := make([]byte, block.BlockSize()+len(data))

	// Generate IV
	iv := encData[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Crypt blocks blocks
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(encData[block.BlockSize():], data)
	return encData, nil
}

func cbcDecrypt(encData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encData) < block.BlockSize() {
		return nil, errors.New("data is too short")
	}

	iv := encData[:block.BlockSize()]

	encData = encData[block.BlockSize():]
	if len(encData)%block.BlockSize() != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encData, encData)

	return pkcsStripPadding(encData), nil
}
