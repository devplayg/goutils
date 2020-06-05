package goutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncAES returns data encrypted with AES
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func EncAES(data, key []byte) ([]byte, error) {
	k := len(key)
	switch k {
	default:
		return nil, aes.KeySizeError(k)
	case 16, 24, 32:
		break
	}
	return cbcEncrypt(data, key)
}

// EncAES256 returns data encrypted with AES-128
func EncAES256(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 32)
}

// EncAES192 returns data encrypted with AES-192
func EncAES192(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 24)
}

// EncAES128 returns data encrypted with AES-256
func EncAES128(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 16)
}

// DecAES256 returns data decrypted with AES-128
func DecAES256(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 32)
}

// DecAES192 returns data decrypted with AES-192
func DecAES192(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 24)
}

// DecAES128 returns data decrypted with AES-256
func DecAES128(data, key []byte) ([]byte, error) {
	return encAESWithKeyLen(data, key, 16)
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
