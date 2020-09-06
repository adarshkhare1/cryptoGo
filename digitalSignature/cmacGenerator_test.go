package digitalSignature

import (
	. "cryptoGo/testHelper"
	"testing"
)

func TestCMAGenerator_GetCheckSumForStringSHA512(t *testing.T) {
	h := NewCMACGenerator(GenerateTestSecret32(t))
	checkSum := h.GeCMACForMessage("This is test")
	LogStatement("String checksum %s\n", checkSum)
}

func TestCMAGenerator_GetCheckSumForFileSHA512(t *testing.T) {
	h := NewCMACGenerator(GenerateTestSecret32(t))
	checkSum  := h.GetHashForFile("mytest.txt")
	LogStatement("File checksum %s\n", checkSum)
}

