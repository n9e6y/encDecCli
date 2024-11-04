package cipher

import "crypto/rsa"

// Base interface for cipher operations
type Cipher interface {
	Encrypt(data []byte, key *rsa.PublicKey) ([]byte, error)
	Decrypt(data []byte, key *rsa.PrivateKey) ([]byte, error)
}
