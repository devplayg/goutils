package goutils

//func TestAes256Cbc2_test(t *testing.T) {
//	type args struct {
//		cipherText string
//		plainText string
//		key        string
//	}
//
//	tests := []struct {
//		name string
//		args args
//	}{
//		{name: "AES-256 decrypt", args: struct {
//			cipherText string
//			plainText  string
//			key        string
//			//}{cipherText: "5XgDF86p1Qr4A4a7al7KQGoLa+N/sw9ySw/Lsd+kOCc36Ma/Dx3zeBFSSQ==", plainText: "no woman no cry", key: "bob marley"}}, // GCM
//		}{cipherText: "hAcrtXPPAL0EI0+9DgxS8JUfZRFvNknuw5YBrZaRQus=", plainText: "no woman no cry", key: "bob marley"}},// CBC
//	}
//
//	for _, tt := range tests {
//		aesCbcUtils := NewAes256CBCUtils([]byte(tt.args.key))
//		decrypted, err := aesCbcUtils.DecryptString(tt.args.cipherText, tt.args.key)
//		//aesGcmUtils := NewAes256GCM([]byte(tt.args.key))
//		//decrypted, err := aesGcmUtils.DecryptString(tt.args.cipherText)
//		if err != nil {
//			t.Error("failed to decrypt")
//		}
//		if decrypted != tt.args.plainText {
//			t.Errorf("Decrypt() = %v, want = %v", decrypted, tt.args.plainText)
//		}
//	}
//}
