package mac

import (
	. "crypto"
	. "cryptoGo/testHelper"
	"testing"
)

func TestHMAGenerator_GetCheckSumForStringSHA512(t *testing.T) {
	h := NewHMACGenerator(SHA512, GenerateTestSecret32(t))
	checkSum := h.GeHMACForMessage("This is test")
	LogStatement("String checksum %s\n", checkSum)
}



func TestHMAGenerator_GetCheckSumForFileSHA512(t *testing.T) {
	h := NewHMACGenerator(SHA512, GenerateTestSecret32(t))
	checkSum  := h.GetHashForFile("mytest.txt")
	LogStatement("File checksum %s\n", checkSum)
}

func TestHMAGenerator_GetCheckSumForStringDefault(t *testing.T) {
	h := NewDefaultHMACGenerator(GenerateTestSecret32(t))
	checkSum := h.GeHMACForMessage("This is test")
	LogStatement("String checksum %s\n", checkSum)
}



func TestHMAGenerator_GetCheckSumForFileDefault(t *testing.T) {
	h := NewDefaultHMACGenerator(GenerateTestSecret32(t))
	checkSum  := h.GetHashForFile("mytest.txt")
	LogStatement("File checksum %s\n", checkSum)
}