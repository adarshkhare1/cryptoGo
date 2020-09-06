package digitalSignature

import (
	. "crypto"
	"crypto/hmac"
	"cryptoGo/hashcode"
	"encoding/hex"
	"io"
	"log"
	"os"
)

type HMACGenerator struct {
	hash Hash
	key []byte
}

func NewDefaultHMACGenerator(secret []byte) *HMACGenerator {
	return NewHMACGenerator(SHA256, secret)
}

func NewHMACGenerator(hash Hash, secret []byte) *HMACGenerator {
	return &HMACGenerator{hash: hash, key: secret}
}


//Compute the HMAC of given message string, used the key and hash algorithm defined for the generator
func (g *HMACGenerator) GeHMACForMessage(data string) string {
	hashFunc := hashcode.GetHashingFunction(g.hash)
	h := hmac.New(hashFunc, g.key)
	h.Write([]byte(data))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

//Compute the HMAC of given file, used the key and hash algorithm defined for the generator
func (g *HMACGenerator) GetHashForFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		defer f.Close()
		hashFunc := hashcode.GetHashingFunction(g.hash)
		h := hmac.New(hashFunc, g.key)
		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}
		result := hex.EncodeToString(h.Sum(nil))
		return result
		}
}

