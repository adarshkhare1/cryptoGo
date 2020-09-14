package asymmetricKey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"cryptoGo/testHelper"
	"encoding/pem"
)

type RSAKeyPair struct {
	PublicKey []byte
	PrivateKey []byte
}

func NewRSAKeyPair(bitSize int) *RSAKeyPair {
	return GenerateRSAKeyPair(bitSize)
}

//
func GenerateRSAKeyPair(bitSize int)  *RSAKeyPair {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		panic(err)
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	encodePrivateKeyToPEM(privateKey)
	testHelper.LogStatement("Private Key generated")
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	testHelper.LogStatement("Public key generated")
	return  &RSAKeyPair{PublicKey: publicKeyBytes, PrivateKey: privateKeyBytes}
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}