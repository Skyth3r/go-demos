package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	inputText := os.Args[2]
	passphrase := os.Args[3]
	encryptFlag := flag.Bool("e", false, "Encrypts the message using the key.")
	decryptFlag := flag.Bool("d", false, "Decrypts the message using the key.")
	flag.Parse()

	if *encryptFlag == true {
		encryptedMessage := encryptMessage(inputText, passphrase)
		fmt.Printf("Encrypted message: %x\n", encryptedMessage)
	} else if *decryptFlag == true {
		decryptedMessage := decryptMessage(inputText, passphrase)
		fmt.Printf("Decrypted message: %s\n", decryptedMessage)
	} else {
		fmt.Println("Please try again")
	}
}

func hashPass(pass string) string {
	hashedpass := sha256.Sum256([]byte(pass))
	hashedBytes := hashedpass[:]
	return string(hashedBytes)
}

func encryptMessage(message string, pass string) string {
	encrpytionKey := hashPass(pass)

	cipherBlock, err := aes.NewCipher([]byte(encrpytionKey))
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatal(err)
	}

	non := make([]byte, aesgcm.NonceSize())
	io.ReadFull(rand.Reader, non)

	cipherText := aesgcm.Seal(non, non, []byte(message), nil)
	return string(cipherText)
}

func decryptMessage(message string, pass string) string {
	encrpytionKey := hashPass(pass)
	encodedMessage, err := hex.DecodeString(message)
	if err != nil {
		log.Fatal(err)
	}

	cipherBlock, err := aes.NewCipher([]byte(encrpytionKey))
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatal(err)
	}

	nonSize := aesgcm.NonceSize()
	non, cipherText := encodedMessage[:nonSize], encodedMessage[nonSize:]

	plainText, err := aesgcm.Open(nil, non, cipherText, nil)
	if err != nil {
		log.Fatal(err)
	}
	return string(plainText)
}
