package crypto

import (
	"testing"
)

func TestEncAes(t *testing.T) {
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
		}{key: []byte("16-BYTES-KEY-###"), data: []byte("no woman no cry~")}},
		{name: "AES-192 encryption", args: struct {
			key  []byte
			data []byte
		}{key: []byte("24-BYTES-KEY-###########"), data: []byte("no woman no cry~")}},
		{name: "AES-256 encryption", args: struct {
			key  []byte
			data []byte
		}{key: []byte("32-BYTES-KEY-###################"), data: []byte("no woman no cry~")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EncAES(tt.args.data, tt.args.key)
		})
	}
}
