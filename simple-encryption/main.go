package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func main() {

	plainText := "This is a message to be encrypted"
	passphrase := "This is a passphrase!!!!"

	encryptedMessage := encryptMessage(plainText, passphrase)
	//decryptedMessage := decryptMessage()

	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	//fmt.Printf("decrypted message: %s\n", decryptedMessage)

}

func encryptMessage(message string, pass string) string {
	cipherBlock, err := aes.NewCipher([]byte(pass))
	if err != nil {
		log.Fatal(err)
	}
	cipherText := make([]byte, len(message))
	cipherBlock.Encrypt(cipherText, []byte(message))
	return string(cipherText)
}

// func decryptMessage(message string, pass string) string {

// }
