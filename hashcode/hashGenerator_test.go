package hashcode

import (
	"fmt"
	"testing"
)

func TestHashGenerator_GetCheckSumForString(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForString("This is test")
	fmt.Printf("String checksum %x\n", checkSum)
}

func TestHashGenerator_GetCheckSumForFile(t *testing.T) {
	h := NewDefaultHashGenerator()
	checkSum, _  := h.GetCheckSumForFile("mytest.txt")
	fmt.Printf("File checksum %x\n", checkSum)
}