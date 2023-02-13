package main

import (
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	plainText := "This is a message to be encrypted"
	passphrase := "Test passphrase!"

	encryptedMessage := encryptMessage(plainText, passphrase)

	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
}

func encryptMessage(message string, pass string) string {
	hashedPass := sha256.Sum256([]byte(pass))
	hashedPassByte := hashedPass[:]
	cipherBlock, err := aes.NewCipher([]byte(hashedPassByte))
	if err != nil {
		log.Fatal(err)
	}
	cipherText := make([]byte, len(message))
	cipherBlock.Encrypt(cipherText, []byte(message))
	return string(cipherText)
}
