package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

//// RSA cipher implementation
//type RSACipher struct{}
//
//func (r *RSACipher) Encrypt(data []byte) ([]byte, error) {
//    fmt.Println("Encrypting with RSA")
//    return data, nil
//}
//
//func (r *RSACipher) Decrypt(data []byte) ([]byte, error) {
//    fmt.Println("Decrypting with RSA")
//    return data, nil
//}

func EncryptRSA(data []byte, key *rsa.PrivateKey) ([]byte, error) {
	if key == nil {
		return nil, errors.New("private key is nil")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, &key.PublicKey, data)
}

func DecryptRSA(ciphertext []byte, key *rsa.PrivateKey) ([]byte, error) {
	if key == nil {
		return nil, errors.New("private key is nil")
	}
	return rsa.DecryptPKCS1v15(rand.Reader, key, ciphertext)

}
