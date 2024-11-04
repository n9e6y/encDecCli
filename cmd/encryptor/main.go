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

	fmt.Println("Operation completed successfully.")
}
