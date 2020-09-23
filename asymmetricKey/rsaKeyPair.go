package asymmetricKey

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"cryptoGo/testHelper"
	"encoding/pem"
	"errors"
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
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	testHelper.LogStatement("Private Key generated\n")
	publicKeyBytes := encodePublicKeyToPEM(&privateKey.PublicKey)

	testHelper.LogStatement("Public key generated\n")
	return  &RSAKeyPair{PublicKey: publicKeyBytes, PrivateKey: privateKeyBytes}
}

//Encrypt a message key with given public key.
func EncryptMessage(message string, key []byte) []byte  {
	block, _ := pem.Decode(key)
	if block == nil {
		panic(errors.New("public key error"))
	}
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}


	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message),
		nil)
	if err != nil {
		panic(err)
	}
	return encryptedBytes
}

//Decrypt message bytes using a given Private Key
func DecryptMessage(encryptedBytes []byte, key []byte) string {
	block, _ := pem.Decode(key)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	decryptedBytes, err := privateKey.Decrypt(nil,
		encryptedBytes,
		&rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
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

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePublicKeyToPEM(publicKey *rsa.PublicKey) []byte {
	// Get ASN.1 DER format
	pubDER := x509.MarshalPKCS1PublicKey(publicKey)

	// pem.Block
	var pubBlock = pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	}
	// Public key in PEM format
	publicPEM := pem.EncodeToMemory(&pubBlock)

	return publicPEM
}