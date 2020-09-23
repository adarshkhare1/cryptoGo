# cryptoGo
Catalog of cryptography golang code samples for various cryptographic algorithms applications.

---
## Symmetric Key Cryptography
Symmetric key cryptography schemes are based on single-key. Applications of symmetric key cryptography includes volume encryption using block ciphers (AES, 3DES), wireless protocols using stream ciphers, preserving message integrity using hashcodes (MD5, SHA1, SHA256, SHA512) , message authenticity and integrity using Message Authentication Codes (MAC).

### Block Encryption (AES, AES-GCM)
Block Cipher process messages block by block, that means an entire block of data is encrypted with the same key.

 - [CFB Mode AES encryption](/blockEncryption/cfbBlockCipher.go) - Encrypt a file using AES encryption. 
 - [Authenticated Cipher Block Encryption](/blockEncryption/authenticatedBlockCipher.go) - AES-GCM Encryption of a file.

### Hashing (MD5, SHA1, SHA256, SHA512)
Hash function compute a digest of a message, which is a short, fixed-length bit array. Hash can be considered as fingerprint of the message that is unique for a given message. Hash function applications are public key encryption, digital signature, message integrity verification, and password protection. 

 -  [Hashcode](/hashcode/hashGenerator.go) - Generate Hashcode for a text file.

### Message Authentication Code (MAC)
Message Authentication code (MAC), also known as cryptographic checksum or keyed hash function. They provide message integrity and message authentication. Mac can be generated either by keyed hashing or cipher block hashing.

 -  [Hash based MAC (HMAC)](/mac/hmacGenerator.go) - Keyed hashing with selected hashing algorithm .
 -  [Cipher Based MAC (CMAC)](/mac/cmacGenerator.go) - Cipher based hashing.
 
## Asymmetric Key Cryptography
Asymmetric Key Cryptography is public key cryptography. It uses a pair of keys, a public key that can be used by anyone to encrypt the message for the key owner. A private key that is only known to the owner and can be used to decrypt the message that is encrypted by the public key.
### RSA Keypair
 - [Generate RSA Key pair as byte arrays](/asymmetricKey/rsaKeyPair.go)
 - [Encrypt / Decrypt a key using RSA key pair](/asymmetricKey/rsaKeyPair.go)