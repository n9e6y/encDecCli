package keyutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// GenerateKeys generates a new RSA private and public key pair and saves them to files.
func GenerateKeys(privateKeyPath, publicKeyPath string) error {
	// Generate RSA keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Save private key
	privFile, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	defer privFile.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	err = pem.Encode(privFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})
	if err != nil {
		return err
	}

	// Save public key
	pubFile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	defer pubFile.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	err = pem.Encode(pubFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})
	if err != nil {
		return err
	}

	return nil
}

// LoadRSAKey
func LoadRSAKey(path, operation string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}
	switch operation {
	case "encrypt":
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaPub, ok := pub.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("failed to parse RSA public key")
		}
		return &rsa.PrivateKey{PublicKey: *rsaPub}, nil
	case "decrypt":
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return priv, nil
	default:
		return nil, errors.New("unknown operation")
	}
}
