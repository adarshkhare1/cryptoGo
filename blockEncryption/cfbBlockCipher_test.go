package blockEncryption

import (
	. "cryptoGo/testHelper"
	"fmt"
	"testing"
)

func TestEncryptFileUsingCFB(t *testing.T) {
	cfbCipher := NewCFBBlockCipher(GenerateTestSecret32(t))
	ciphertext, err := cfbCipher.EncryptFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	LogStatement("ENCRYPT\n%x\n", ciphertext)
	plainText, err := cfbCipher.Decrypt(ciphertext)
	if err != nil {
		LogError(err)
	}
	LogStatement("DECRYPT\n%s\n", string(plainText))
}