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

//Compute the Hash of given data string using the SHA256 checksum.
func (g *HashGenerator) GetCheckSumForString(data string) ([]byte, error) {
	h, err := instantiateHashFunction(g.hash)
	if err == nil {
		return g.getHashForString(h, data), nil
	} else {
		return nil, err
	}
}

//Compute the Hash of given data string using the SHA256 checksum.
func (g *HashGenerator) GetCheckSumForFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		defer f.Close()
		h, err := instantiateHashFunction(g.hash)
		if err == nil {
			return g.getHashForFile(h, f), nil
		} else {
			return nil, err
		}
	}
}

	//Compute the Hash of given data string using the SHA256 checksum.
func (g *HashGenerator) getHashForString(h hash.Hash, data string) []byte {
	h.Write([]byte(data))
	return h.Sum(nil)
}

//Compute the Hash of given file using the SHA256 checksum
func (g *HashGenerator) getHashForFile(h hash.Hash, f *os.File) []byte {
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		return nil
	}
	return h.Sum(nil)
}

//instantiateHashFunction instantiate a Hash generator for a given algorithm.
func instantiateHashFunction(algorithm Hash) (hash.Hash, error) {
	switch algorithm {
	case MD5:
		return md5.New(), nil
	case SHA1:
		return sha1.New(), nil
	case SHA256:
		return sha256.New(), nil
	case SHA512:
		return sha512.New(), nil
	default:
		return nil, errors.New("algorithm not supported")
	}
}