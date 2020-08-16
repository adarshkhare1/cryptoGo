package hashcode

import "strconv"

// Hash identifies a cryptographic hash function that is implemented in another
// package.
type Hash uint

// HashFunc simply returns the value of h so that Hash implements SignerOpts.
func (h Hash) HashFunc() Hash {
	return h
}


func (h Hash) String() string {
	switch h {
	case MD4:
		return "MD4"
	case MD5:
		return "MD5"
	case SHA1:
		return "SHA-1"
	case SHA224:
		return "SHA-224"
	case SHA256:
		return "SHA-256"
	case SHA384:
		return "SHA-384"
	case SHA512:
		return "SHA-512"
	case MD5SHA1:
		return "MD5+SHA1"
	case RIPEMD160:
		return "RIPEMD-160"
	case SHA3_224:
		return "SHA3-224"
	case SHA3_256:
		return "SHA3-256"
	case SHA3_384:
		return "SHA3-384"
	case SHA3_512:
		return "SHA3-512"
	case SHA512_224:
		return "SHA-512/224"
	case SHA512_256:
		return "SHA-512/256"
	case BLAKE2s_256:
		return "BLAKE2s-256"
	case BLAKE2b_256:
		return "BLAKE2b-256"
	case BLAKE2b_384:
		return "BLAKE2b-384"
	case BLAKE2b_512:
		return "BLAKE2b-512"
	default:
		return "unknown hash value " + strconv.Itoa(int(h))
	}

}

const (

	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b
	maxHash
)
