package blockEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func EncryptFile(fileName string, key []byte) ([]byte, error){
	//read the file as cyphertext
	plainText, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//create new aes cypher
	block, err := aes.NewCipher(key)
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
	ciphertext := aesgcm.Seal(nonce, nonce, plainText, nil)
	return ciphertext, nil
}

func Decrypt(cipherText []byte, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
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