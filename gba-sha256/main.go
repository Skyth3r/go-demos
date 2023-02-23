package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	text := "This is a string to be hashed!"
	hashedText := sha256.New()
	hashedText.Write([]byte(text))

	byteSlice := hashedText.Sum(nil)

	fmt.Printf("%x\n", byteSlice)
}
