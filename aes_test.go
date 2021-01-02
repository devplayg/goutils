package goutils

import (
	"testing"
)

func AesKeyTest(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "AES-256-Key", args: struct {
			key []byte
		}{
			key: []byte("32-BYTES-KEY-###################"),
		}},
	}

	for _, tt := range tests {
		if key := Aes256Key(tt.args.key); len(key) != 16 {
			t.Errorf("Aes256Key() = %v, want %v", len(key), 32)
		}

		if key := Aes128Key(tt.args.key); len(key) != 16 {
			t.Errorf("Aes256Key() = %v, want %v", len(key), 16)
		}
	}
}
