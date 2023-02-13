package main

import (
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	plainText := "This is a message to be encrypted & decrypted"
	passphrase := "Test passphrase!"

	encryptedMessage := encryptMessage(plainText, passphrase)
	decryptedMessage := decryptMessage(encryptedMessage, passphrase)

	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)
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
func decryptMessage(message string, pass string) string {
	hashedPass := sha256.Sum256([]byte(pass))
	hashedPassByte := hashedPass[:]
	cipherBlock, err := aes.NewCipher([]byte(hashedPassByte))
	if err != nil {
		log.Fatal(err)
	}
	cipherText := make([]byte, len(message))
	cipherBlock.Decrypt(cipherText, []byte(message))
	return string(cipherText)
}
