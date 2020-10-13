package utility

import (
	"testing"
)

var key = "thisisasecretkeystring"

func TestAESEncrypterWithCFB(t *testing.T) {
	//key := make([]byte, 8)
	//_, err := rand.Read(key)
	//if err != nil {
	//	t.Error(err)
	//}
	keyBytes, err := ParseKey(key)
	if err != nil {
		t.Error(err)
	}

	cipherText, err := AESEncrypterWithCFB("hello", keyBytes)
	if err != nil {
		t.Error(err)
	}
	t.Log(cipherText)

	plainText, err := AESDecrypterWithCFB(cipherText, keyBytes)
	if err != nil {
		t.Error(err)
	}
	t.Log(plainText)
}
