package hashcode

import (
	. "crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
	"log"
	"os"
)


type HashGenerator struct {
	hash Hash
}

func NewHashGenerator(hash Hash) *HashGenerator {
	return &HashGenerator{hash: hash}
}

func NewDefaultHashGenerator() *HashGenerator {
	return NewHashGenerator(SHA256)
}

//Compute the Hash of given data string, using the algorithm specified for generator.
func (g *HashGenerator) GetCheckSumForString(data string) ([]byte, error) {
	if data == "" {
		return nil, errors.New("empty message to hash")
	}
	h := GetHashingFunction(g.hash)()
	h.Write([]byte(data))
	return h.Sum(nil), nil
}

//Compute the Hash of given data string, using the algorithm specified for generator.
func (g *HashGenerator) GetCheckSumForFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
		h := GetHashingFunction(g.hash)()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
			return nil, err
		}
		return h.Sum(nil), nil
	}
}

//getHashingFunction instantiate a Hash generator for a given algorithm.
func GetHashingFunction(algorithm Hash) func() hash.Hash {
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