package digitalSignature

import (
	"encoding/hex"
	"fmt"
	"github.com/jacobsa/crypto/cmac"
	"io"
	"os"
)

type CMACGenerator struct {
	key []byte
}

func NewCMACGenerator(secret []byte) *CMACGenerator {
	return &CMACGenerator{key: secret}
}


//Compute the HMAC of given message string, used the key and hash algorithm defined for the generator
func (g *CMACGenerator) GeCMACForMessage(data string) string {
	c, err := cmac.New(g.key)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize CMAC: %s", err.Error()))
	}
	c.Write([]byte(data))
	result := hex.EncodeToString(c.Sum(nil))
	return result
}

//Compute the HMAC of given file, used the key and hash algorithm defined for the generator
func (g *CMACGenerator) GetHashForFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	} else {
		defer f.Close()
		c, err := cmac.New(g.key)
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize CMAC: %s", err.Error()))
		}
		if _, err := io.Copy(c, f); err != nil {
			panic(err)
		}
		result := hex.EncodeToString(c.Sum(nil))
		return result
	}
}



