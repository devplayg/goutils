package goutils

import (
	"bytes"
	"testing"
)

var nonce = key(12)

func TestEncAesGcm(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
		nonce []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-128 encryption", args: struct {
			key  []byte
			data []byte
			nonce []byte
		}{key: key16B, data: plainText, nonce: nonce}},
		{name: "AES-256 encryption", args: struct {
			key  []byte
			data []byte
			nonce []byte
		}{key: key16B, data: plainText, nonce: nonce}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := EncAesGcm(tt.args.data, tt.args.key, tt.args.nonce)
			if err != nil {
				t.Errorf("EncAES(); %v", err.Error())
			}

			decrypted, err := DecAesGcm(encrypted, tt.args.key, tt.args.nonce)
			if err != nil {
				t.Errorf("DecAES(); %v", err.Error())
			}

			if !bytes.Equal(decrypted, plainText) {
				t.Errorf("decrypt() = %v, want %v", decrypted, tt.args)
			}
		})
	}
}