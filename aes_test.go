package goutils

import (
	"testing"
)

func TestAes256KeyTest(t *testing.T) {
	type args struct {
		key    []byte
		keyLen int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-256-Key", args: struct {
			key    []byte
			keyLen int
		}{
			key:    []byte("32-byte key"),
			keyLen: 32,
		}},
	}

	for _, tt := range tests {
		if key := Aes256Key(tt.args.key); len(key) != tt.args.keyLen {
			t.Errorf("Aes256Key() = %v, want %v", len(key), tt.args.keyLen)
		}
	}
}

func TestAes128KeyTest(t *testing.T) {
	type args struct {
		key    []byte
		keyLen int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-128-Key", args: struct {
			key    []byte
			keyLen int
		}{
			key:    []byte("16-byte key"),
			keyLen: 16,
		}},
	}

	for _, tt := range tests {
		if key := Aes128Key(tt.args.key); len(key) != 16 {
			t.Errorf("Aes128Key() = %v, want %v", len(key), tt.args.keyLen)
		}
	}
}
