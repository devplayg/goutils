package goutils

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func decodeString(str string) []byte {
	b, _ := hex.DecodeString(str)
	return b
}

func TestBytesToInt64(t *testing.T) {
	type args struct {
		n int64
		b []byte
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "num", args: args{
			n: 0,
			b: decodeString("0000000000000000"),
		}},
		{name: "num", args: args{
			n: 92233720368547758,
			b: decodeString("0147ae147ae147ae"),
		}},
		{name: "num", args: args{
			n: 184467440737095516,
			b: decodeString("028f5c28f5c28f5c"),
		}},
		{name: "num", args: args{
			n: 1844674407370955160,
			b: decodeString("1999999999999998"),
		}},
	}
	for _, tt := range tests {
		var result1 = BytesToInt64(tt.args.b)
		if result1 != tt.args.n {
			t.Errorf("result %v, want %v", result1, tt.args.n)
		}

		var result2 = Int64ToBytes(tt.args.n)
		if !bytes.Equal(result2, tt.args.b) {
			t.Errorf("result %v, want %v", result2, tt.args.n)
		}

	}

	/*
		92233720368547758	0147ae147ae147ae
		184467440737095516	028f5c28f5c28f5c
		276701161105643274	03d70a3d70a3d70a
		368934881474191032	051eb851eb851eb8
		461168601842738790	0666666666666666
		553402322211286548	07ae147ae147ae14
		645636042579834306	08f5c28f5c28f5c2
		737869762948382064	0a3d70a3d70a3d70
		830103483316929822	0b851eb851eb851e
		922337203685477580	0ccccccccccccccc
		1014570924054025338	0e147ae147ae147a
		1106804644422573096	0f5c28f5c28f5c28
		1199038364791120854	10a3d70a3d70a3d6
		1291272085159668612	11eb851eb851eb84
		1383505805528216370	1333333333333332
		1475739525896764128	147ae147ae147ae0
		1567973246265311886	15c28f5c28f5c28e
		1660206966633859644	170a3d70a3d70a3c
		1752440687002407402	1851eb851eb851ea
		1844674407370955160	1999999999999998
		--- PASS: TestBytesToInt64 (0.00s)
		PASS

		Process finished with exit code 0

	*/
	//var i int64
	//var count = 0
	//for i = 0; i < math.MaxInt64; i = i + (math.MaxInt64 / 100) {
	//	//fmt.Println("{name: \"num\", args: args{")
	//	fmt.Printf("n: %d,\nb: decodeString("%x,\n", i, Int64ToBytes(i))
	//	fmt.Println("}},")
	//	count++
	//	if count > 20 {
	//		return
	//	}
	//}

}
