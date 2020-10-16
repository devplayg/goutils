package goutils

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"testing"
)

var (
	plainText = []byte("No woman no cry~")
	key16B    = key(16)
	key24B    = key(24)
	key32B    = key(32)
)

func key(size int) []byte {
	key := make([]byte, size)
	_, _ = rand.Read(key)
	return key
}

func TestEncAesCbc(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-128 encryption", args: struct {
			key  []byte
			data []byte
		}{key: key16B, data: plainText}},
		{name: "AES-192 encryption", args: struct {
			key  []byte
			data []byte
		}{key: key24B, data: plainText}},
		{name: "AES-256 encryption", args: struct {
			key  []byte
			data []byte
		}{key: key32B, data: plainText}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := EncAesCbc(tt.args.data, tt.args.key)
			if err != nil {
				t.Errorf("EncAES(); %v", err.Error())
			}

			decrypted, err := DecAesCbc(encrypted, tt.args.key)
			if err != nil {
				t.Errorf("DecAES(); %v", err.Error())
			}

			if !bytes.Equal(decrypted, plainText) {
				t.Errorf("decrypt() = %v, want %v", decrypted, tt.args)
			}


			k := len(tt.args.key)
			switch k {
			case 16:
				encrypted, err := EncAes128Cbc(tt.args.data, tt.args.key)
				if err != nil {
					t.Errorf("EncAes128Cbc(); %v", err.Error())
				}

				decrypted, err := DecAes128Cbc(encrypted, tt.args.key)
				if err != nil {
					t.Errorf("DecAes128Cbc(); %v", err.Error())
				}

				if !bytes.Equal(decrypted, plainText) {
					t.Errorf("decrypt() = %v, want %v", decrypted, tt.args)
				}
				break

			case 24:
				encrypted, err := EncAes192Cbc(tt.args.data, tt.args.key)
				if err != nil {
					t.Errorf("EncAes128Cbc(); %v", err.Error())
				}

				decrypted, err := DecAes192Cbc(encrypted, tt.args.key)
				if err != nil {
					t.Errorf("DecAes128Cbc(); %v", err.Error())
				}

				if !bytes.Equal(decrypted, plainText) {
					t.Errorf("decrypt() = %v, want %v", decrypted, tt.args)
				}
				break

			case 32:
				encrypted, err := EncAes256Cbc(tt.args.data, tt.args.key)
				if err != nil {
					t.Errorf("EncAes256Cbc(); %v", err.Error())
				}

				decrypted, err := DecAes256Cbc(encrypted, tt.args.key)
				if err != nil {
					t.Errorf("DecAes256Cbc(); %v", err.Error())
				}

				if !bytes.Equal(decrypted, plainText) {
					t.Errorf("decrypt() = %v, want %v", decrypted, tt.args)
				}
				break

			default:
				t.Errorf("%v", aes.KeySizeError(k))
			}
		})
	}
}