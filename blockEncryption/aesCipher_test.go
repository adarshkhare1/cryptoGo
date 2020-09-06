package blockEncryption

import (
	"crypto/rand"
	. "cryptoGo"
	"fmt"
	"testing"
)

func TestEncryptFile(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Error("fail to generate key")
	}
	ciphertext, err := EncryptFile("test.txt", key)
	if err != nil {
		fmt.Println(err)
	}
	LogStatement("%x\n", ciphertext)
    plainText, err := Decrypt(ciphertext, key)
	if err != nil {
		LogError(err)
	}
	LogStatement("%s\n", string(plainText))
}
