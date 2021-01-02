package goutils

import (
	"bytes"
	"testing"
)

func TestAesGCM(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-256-GCM encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  []byte("32-BYTES-KEY-###################"),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-128-GCM encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  []byte("16-BYTES-KEY-###"),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-256-GCM encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  Aes256Key([]byte("ANY VALUE")),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},

		{name: "AES-128-GCM encryption", args: struct {
			key  []byte
			data []byte
		}{
			key:  Aes128Key([]byte("ANY VALUE")),
			data: []byte("no woman no cry~by-레게이즈더뮤직"),
		}},
	}

	for _, tt := range tests {
		aesGcmUtils := NewAesGCM(tt.args.key)
		encrypted, err := aesGcmUtils.Encrypt(tt.args.data)
		if err != nil {
			t.Errorf("Encrypt(); %v", err.Error())
		}
		decrypted, err := aesGcmUtils.Decrypt(encrypted)
		if err != nil {
			t.Errorf("Decrypt(); %v", err.Error())
		}
		if !bytes.Equal(tt.args.data, decrypted) {
			t.Errorf("Decrypt() = %v, want %v", decrypted, tt.args.data)
		}
	}
}
