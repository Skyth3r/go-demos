package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func main() {
	plainText := "This is a message to be encrypted & decrypted"
	passphrase := "Test passphrase!"

	encryptedMessage := encryptMessage(plainText, passphrase)
	decryptedMessage := decryptMessage(encryptedMessage, passphrase)

	fmt.Printf("Plain Text: %s\n", plainText)
	fmt.Printf("Encrypted message: %x\n", encryptedMessage)
	fmt.Printf("Decrypted message: %x\n", decryptedMessage)
}

func hashPass(pass string) string {
	sha256Pass := sha256.Sum256([]byte(pass))
	hashedPassByte := sha256Pass[:]
	return string(hashedPassByte)
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
	if _, err := io.ReadFull(rand.Reader, non); err != nil {
		log.Fatal(err)
	}

	cipherText := aesgcm.Seal(nil, non, []byte(message), nil)
	return string(cipherText)
}

func decryptMessage(message string, pass string) string {
	byteMessage := []byte(message)
	encrpytionKey := hashPass(pass)
	cipherBlock, err := aes.NewCipher([]byte(encrpytionKey))
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatal(err)
	}

	nonSize := aesgcm.NonceSize()
	non, cipherText := byteMessage[:nonSize], byteMessage[nonSize:]

	plainText, err := aesgcm.Open(nil, non, cipherText, nil)
	if err != nil {
		log.Fatal(err)
	}
	return string(plainText)
}
