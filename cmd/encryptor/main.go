package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"encDecCli/pkg/cipher"
	"encDecCli/pkg/fileutil"
	"encDecCli/pkg/keyutil"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("in", "", "Path to the input file")
	outputFile := flag.String("out", "", "Path to the output file")
	operation := flag.String("op", "encrypt", "Operation to perform: encrypt or decrypt")
	keyFile := flag.String("key", "", "Path to the RSA key file")
	keygen := flag.Bool("keygen", false, "Generate RSA keys")
	flag.Parse()

	// Validate required flags
	if *inputFile == "" || *outputFile == "" || *keyFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *keygen {
		err := keyutil.GenerateKeys("private.pem", "public.pem")
		if err != nil {
			log.Fatalf("Key generation failed: %v", err)
		}
		fmt.Println("RSA keys generated successfully.")
		os.Exit(0)
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

	//fmt.Println("successfully generated key")
	fmt.Println("Operation completed successfully.")
}
