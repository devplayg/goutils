package compress

import (
	"bytes"
	"testing"
)

func TestGzip(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "Gzip", args: struct{ data []byte }{data: []byte("devplayg")}, want: []byte("devplayg")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zipped, err := Gzip(tt.args.data)
			if err != nil {
				t.Errorf("Gzip() error = %v", err)
				return
			}
			unzipped, err := Gunzip(zipped)
			if err != nil {
				t.Errorf("Gunzip() error = %v", err)
				return
			}
			if !bytes.Equal(tt.args.data, unzipped) {
				t.Errorf("data mismatched")
				return
			}
		})
	}
}
