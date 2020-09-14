package asymmetricKey

import (
	"cryptoGo/testHelper"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	bitSize := 4096
	keyPair := NewRSAKeyPair(bitSize)
	///Todo: Add actual test validation to validate keys.
	testHelper.LogStatement("Private Key = \n%x\n",keyPair.PrivateKey)
	testHelper.LogStatement("Public Key = \n%x\n",keyPair.PublicKey)
}