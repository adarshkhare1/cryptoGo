package digitalSignature

import (
	. "crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
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
	hashFunc := getHashingFunction(g.hash)
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
		hashFunc := getHashingFunction(g.hash)
		h := hmac.New(hashFunc, g.key)
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
			panic(err)
		}
		result := hex.EncodeToString(h.Sum(nil))
		return result
		}
}

//getHashingFunction instantiate a Hash generator for a given algorithm.
func getHashingFunction(algorithm Hash) func() hash.Hash {
	switch algorithm {
	case MD5:
		return md5.New
	case SHA1:
		return sha1.New
	case SHA256:
		return sha256.New
	case SHA512:
		return sha512.New
	default:
		panic("algorithm not supported.")
	}
}
