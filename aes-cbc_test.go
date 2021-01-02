package goutils

import (
	"bytes"
	"testing"
)

func TestAesCBC(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-256-CBC encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  []byte("32-BYTES-KEY-###################"),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-128-CBC encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  []byte("16-BYTES-KEY-###"),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-256-CBC encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  Aes256Key([]byte("ANY VALUE")),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-128-CBC encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  Aes128Key([]byte("ANY VALUE")),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},
	}

	for _, tt := range tests {
		aesCbcUtils := NewAesCBC(tt.args.key)
		encrypted, err := aesCbcUtils.Encrypt(tt.args.data)
		if err != nil {
			t.Errorf("Encrypt(); %v", err.Error())
		}
		decrypted, err := aesCbcUtils.Decrypt(encrypted)
		if err != nil {
			t.Errorf("Decrypt(); %v", err.Error())
		}
		if !bytes.Equal(tt.args.data, decrypted) {
			t.Errorf("Decrypt() = %v, want %v", decrypted, tt.args.data)
		}
	}
}

func TestAesCBCString(t *testing.T) {
	type args struct {
		cipherText string
		plainText  string
		key        []byte
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "AES-256 decrypt", args: struct {
			cipherText string
			plainText  string
			key        []byte
		}{
			cipherText: "hAcrtXPPAL0EI0+9DgxS8JUfZRFvNknuw5YBrZaRQus=",
			plainText:  "no woman no cry",
			key:        Aes256Key([]byte("bob marley")),
		}}, // CBC
	}

	for _, tt := range tests {
		aesCBC := NewAesCBC(tt.args.key)
		decrypted, err := aesCBC.DecryptString(tt.args.cipherText)
		if err != nil {
			t.Error("failed to decrypt")
		}
		if decrypted != tt.args.plainText {
			t.Errorf("Decrypt() = %v, want = %v", decrypted, tt.args.plainText)
		}
	}
}
