package hashcode

import (
	. "crypto"
	. "cryptoGo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashGenerator_GetCheckSumForStringDefault(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForString("This is test")
	LogStatement("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileDefault(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	LogStatement("File checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForUnknownFile(t *testing.T) {
	h := NewDefaultHashGenerator()
	_, err  := h.GetCheckSumForFile("NonExistFilet.txt")
	assert.NotNil(t, err, "error is not nil")
}

func TestHashGenerator_GetCheckSumForStringMD5(t *testing.T) {
	h := NewHashGenerator(MD5)
	checkSum, _  := h.GetCheckSumForString("This is test")
	LogStatement("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileMD5(t *testing.T) {
	h := NewHashGenerator(MD5)
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	LogStatement("File checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForStringSHA1(t *testing.T) {
	h := NewHashGenerator(SHA1)
	checkSum, _  := h.GetCheckSumForString("This is test")
	LogStatement("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileSHA1(t *testing.T) {
	h := NewHashGenerator(SHA1)
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	LogStatement("File checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForStringSHA512(t *testing.T) {
	h := NewHashGenerator(SHA512)
	checkSum, _  := h.GetCheckSumForString("This is test")
	LogStatement("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileSHA512(t *testing.T) {
	h := NewHashGenerator(SHA512)
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	LogStatement("File checksum %x\n", checkSum)
}