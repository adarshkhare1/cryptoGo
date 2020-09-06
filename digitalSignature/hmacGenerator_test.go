package digitalSignature

import (
	. "crypto"
	rand "crypto/rand"
	"fmt"
	"testing"
)

func TestHashGenerator_GetCheckSumForStringSHA512(t *testing.T) {
	h := NewHMACGenerator(SHA512, generateTestSecret(t))
	checkSum := h.GeHMACForMessage("This is test")
	fmt.Printf("String checksum %s\n", checkSum)
}



func TestHashGenerator_GetCheckSumForFileSHA512(t *testing.T) {
	h := NewHMACGenerator(SHA512, generateTestSecret(t))
	checkSum  := h.GetHashForFile("mytest.txt")
	fmt.Printf("File checksum %s\n", checkSum)
}

func TestHashGenerator_GetCheckSumForStringDefault(t *testing.T) {
	h := NewDefaultHMACGenerator(generateTestSecret(t))
	checkSum := h.GeHMACForMessage("This is test")
	fmt.Printf("String checksum %s\n", checkSum)
}



func TestHashGenerator_GetCheckSumForFileDefault(t *testing.T) {
	h := NewDefaultHMACGenerator(generateTestSecret(t))
	checkSum  := h.GetHashForFile("mytest.txt")
	fmt.Printf("File checksum %s\n", checkSum)
}


func generateTestSecret(t *testing.T) []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Error("fail to generate key")
	}
	return key
}