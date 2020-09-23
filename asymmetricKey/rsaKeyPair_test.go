package asymmetricKey

import (
	"cryptoGo/testHelper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	bitSize := 4096
	keyPair := NewRSAKeyPair(bitSize)
	///Todo: Add actual test validation to validate keys.
	testHelper.LogStatement("Private Key = \n%x\n",keyPair.PrivateKey)
	testHelper.LogStatement("Public Key = \n%x\n",keyPair.PublicKey)
}

func TestEncryptAndDecryptUsingRSAKeyPair(t *testing.T) {
	bitSize := 4096
	keyPair := NewRSAKeyPair(bitSize)
	message := "This_is_my_code"
	encryptedMessage := EncryptMessage(message, keyPair.PublicKey)
	finalMessage := DecryptMessage(encryptedMessage, keyPair.PrivateKey)
	assert.Equal(t,message, finalMessage)
}