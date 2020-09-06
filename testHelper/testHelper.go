package testHelper

import (
	"crypto/rand"
	"fmt"
	"testing"

)

const LogEnabled = false

func LogStatement(format string, a ...interface{}) (n int, err error){
	if LogEnabled {
		return fmt.Printf(format, a ...)
	}
	return 0, nil
}

func LogError(errToLog error) (n int, err error){
	if LogEnabled {
		return fmt.Println(errToLog)
	}
	return 0, nil
}

func GenerateTestSecret32(t *testing.T) []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Error("fail to generate key")
	}
	LogStatement("Random is %d %x\n", len(key), key)
	return key
}
