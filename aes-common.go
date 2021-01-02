package goutils

import (
	"bytes"
	"crypto/sha256"
)

func Aes256Key(key []byte) []byte {
	key32Bytes := sha256.Sum256(key)
	return key32Bytes[:]
}

func Aes128Key(key []byte) []byte {
	key32Bytes := sha256.Sum256(key)
	return key32Bytes[:16]
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
