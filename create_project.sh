#!/bin/bash

# Define the base directory
base_dir="encDecCli"

# Create the directory structure
mkdir -p $base_dir/cmd/encryptor
mkdir -p $base_dir/pkg/cipher
mkdir -p $base_dir/pkg/fileutil
mkdir -p $base_dir/pkg/keyutil

# Create main.go in cmd/encryptor with package declaration
cat <<EOL > $base_dir/cmd/encryptor/main.go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Encryptor CLI tool")
}
EOL

# Create cipher.go in pkg/cipher with package declaration
cat <<EOL > $base_dir/pkg/cipher/cipher.go
package cipher

import "fmt"

// Base interface for cipher operations
type Cipher interface {
    Encrypt(data []byte) ([]byte, error)
    Decrypt(data []byte) ([]byte, error)
}
EOL

# Create aes_cipher.go in pkg/cipher with package declaration
cat <<EOL > $base_dir/pkg/cipher/aes_cipher.go
package cipher

import "fmt"

// AES cipher implementation
type AESCipher struct{}

func (a *AESCipher) Encrypt(data []byte) ([]byte, error) {
    fmt.Println("Encrypting with AES")
    return data, nil
}

func (a *AESCipher) Decrypt(data []byte) ([]byte, error) {
    fmt.Println("Decrypting with AES")
    return data, nil
}
EOL

# Create rsa_cipher.go in pkg/cipher with package declaration
cat <<EOL > $base_dir/pkg/cipher/rsa_cipher.go
package cipher

import "fmt"

// RSA cipher implementation
type RSACipher struct{}

func (r *RSACipher) Encrypt(data []byte) ([]byte, error) {
    fmt.Println("Encrypting with RSA")
    return data, nil
}

func (r *RSACipher) Decrypt(data []byte) ([]byte, error) {
    fmt.Println("Decrypting with RSA")
    return data, nil
}
EOL

# Create file.go in pkg/fileutil with package declaration
cat <<EOL > $base_dir/pkg/fileutil/file.go
package fileutil

import "fmt"

// Utility functions for file handling
func ReadFile(path string) ([]byte, error) {
    fmt.Println("Reading file:", path)
    return []byte{}, nil
}

func WriteFile(path string, data []byte) error {
    fmt.Println("Writing file:", path)
    return nil
}
EOL

# Create key.go in pkg/keyutil with package declaration
cat <<EOL > $base_dir/pkg/keyutil/key.go
package keyutil

import "fmt"

// Utility functions for key management
func GenerateKey() ([]byte, error) {
    fmt.Println("Generating key")
    return []byte{}, nil
}
EOL

# Create go.mod with module declaration
cat <<EOL > $base_dir/go.mod
module encDecCli

go 1.20
EOL

# Create placeholders for README.md and LICENSE
echo "# encDecCli - File Encryption/Decryption CLI Tool" > $base_dir/README.md
echo "MIT License" > $base_dir/LICENSE

echo "Project structure created successfully!"
