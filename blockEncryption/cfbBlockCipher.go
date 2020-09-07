package blockEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
)

type CFBBlockCipher struct {
	key []byte
}

func NewCFBBlockCipher(secret []byte) *CFBBlockCipher {
	return &CFBBlockCipher{key: secret}
}

func (cfb *CFBBlockCipher) EncryptFile(fileName string) ([]byte, error){
	//read the file as cyphertext
	plainText, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return cfb.EncryptMessage(plainText)
}

func (cfb *CFBBlockCipher) EncryptMessage(message []byte) ([]byte, error) {
	//create new aes cypher
	block, err := aes.NewCipher(cfb.key)
	if err != nil {
		return nil, err
	}
	//Secure AES require an unique IV, adding additional bytes in beginning of the block
	cipherText := make([]byte, aes.BlockSize+len(message))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], message)
	return cipherText, nil
}

func (cfb *CFBBlockCipher) Decrypt(cipherText []byte) ([]byte, error){
	block, err := aes.NewCipher(cfb.key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("invalid ciphertext block")
	}
	plaintext := make([]byte, aes.BlockSize+len(cipherText))
	//An unique IV needs to be in the beginning
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(plaintext, cipherText)
	return plaintext, nil
}