package hashcode

import (
	"fmt"
	"testing"
)

func TestHashGenerator_GetCheckSumForStringDefault(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForString("This is test")
	fmt.Printf("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileDefault(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	fmt.Printf("File checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForStringMD5(t *testing.T) {
	h := NewHashGenerator(MD5)
	checkSum, _  := h.GetCheckSumForString("This is test")
	fmt.Printf("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileMD5(t *testing.T) {
	h := NewHashGenerator(MD5)
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	fmt.Printf("File checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForStringSHA1(t *testing.T) {
	h := NewHashGenerator(SHA1)
	checkSum, _  := h.GetCheckSumForString("This is test")
	fmt.Printf("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFileSHA1(t *testing.T) {
	h := NewHashGenerator(SHA1)
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	fmt.Printf("File checksum %x\n", checkSum)
}