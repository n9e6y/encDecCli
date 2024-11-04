//package main
//
//import (
//    "flag"
//)
//
//func main()
//    // Define flags
//	inputFile := flag.String("input", "", "path to input file")
//	outputFile := flag.String("output", "", "path to output file")
//	operation := flag.String("operation", "encrypt", "operation to encrypt")
//	keyFile := flag.String("key", "", "path to key file")
//	flag.Parse()
//
//    // Validate required flags
//    if *inputFile == "" || *outputFile == "" || *keyFile == "" {
//        flag.Usage()
//        os.Exit(1)
//    }
//
//    // Read the input file
//    data , err := ioutil.ReadFile(*inputFile)
//    if err != nil {
//        log.Fatalf("failed to read the input file" , err)}
//
//    // Load the RSA key
//    rsaKey, err := keyutil.LoadRSAKey(*keyFile , *operation)
//    if err != nil {
//        log.Fatalf("failed to load the key file", err)}
//    var result []byte
//    switch *operation {
//    case "encrypt":
//        result, err = encryptor.Encrypt(data, rsaKey)
//        if err != nil {
//            log.Fatalf("failed to encrypt", err)}}
//    case "decrypt":
//        result, err = decryptor.Decrypt(data, rsaKey)
//        if err != nil {
//            log.Fatalf("failed to decrypt", err)}
//        default:
//            fmt.Printf("unknown operation %s\n", *operation)
//            os.Exit(1)
//
//}
//
//

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"os"

	"encDecCli/pkg/cipher"
	"encDecCli/pkg/fileutil"
	"encDecCli/pkg/keyutil"
)

func generateKeys() {

	// generate RSA keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to open private.pem : %v", err)
	}
	privFile, err := os.Create("private.pem")
	if err != nil {
		log.Fatalf("failed to open private.pem : %v", err)
	}
	defer privFile.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pem.Encode(privFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	// Save public key
	pubFile, err := os.Create("public.pem")
	if err != nil {
		log.Fatalf("failed to open public.pem : %v", err)
	}
	defer pubFile.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("failed to marshal public.pem : %v", err)
	}
	pem.Encode(pubFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	})
}

func main() {
	// Define command-line flags
	inputFile := flag.String("in", "", "Path to the input file")
	outputFile := flag.String("out", "", "Path to the output file")
	operation := flag.String("op", "encrypt", "Operation to perform: encrypt or decrypt")
	keyFile := flag.String("key", "", "Path to the RSA key file")
	flag.Parse()

	// Validate required flags
	if *inputFile == "" || *outputFile == "" || *keyFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Read the input file
	data, err := fileutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Load the RSA key
	rsaKey, err := keyutil.LoadRSAKey(*keyFile, *operation)
	if err != nil {
		log.Fatalf("Failed to load RSA key: %v", err)
	}

	var result []byte
	// Perform the operation
	switch *operation {
	case "encrypt":
		result, err = cipher.EncryptRSA(data, rsaKey)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}
	case "decrypt":
		result, err = cipher.DecryptRSA(data, rsaKey)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}
	default:
		fmt.Println("Invalid operation. Use 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}

	// Write the output file
	err = fileutil.WriteFile(*outputFile, result)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	generateKeys()
	//fmt.Println("successfully generated key")
	fmt.Println("Operation completed successfully.")
}
