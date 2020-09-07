package blockEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

type AuthenticatedBlockCipher struct {
	key []byte
}

func NewAuthenticatedBlockCipher(secret []byte) *AuthenticatedBlockCipher {
	return &AuthenticatedBlockCipher{key: secret}
}

func (ae *AuthenticatedBlockCipher) EncryptFile(fileName string) ([]byte, error){
	//read the file as cyphertext
	plainText, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ae.EncryptMessage(plainText)
}

func (ae *AuthenticatedBlockCipher) EncryptMessage(message []byte) ([]byte, error) {
	//create new aes cypher
	block, err := aes.NewCipher(ae.key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nonce, nonce, message, nil)
	return ciphertext, nil
}

func (ae *AuthenticatedBlockCipher) Decrypt(cipherText []byte) ([]byte, error){
	block, err := aes.NewCipher(ae.key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesgcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}