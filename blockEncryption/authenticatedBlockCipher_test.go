package blockEncryption

import (
	. "cryptoGo/testHelper"
	"fmt"
	"testing"
)

func TestEncryptFile(t *testing.T) {
	ae := NewAuthenticatedBlockCipher(GenerateTestSecret32(t))
	ciphertext, err := ae.EncryptFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	LogStatement("ENCRYPT\n%x\n", ciphertext)
    plainText, err := ae.Decrypt(ciphertext)
	if err != nil {
		LogError(err)
	}
	LogStatement("DECRYPT\n%s\n", string(plainText))
}
