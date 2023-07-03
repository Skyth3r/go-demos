package main

import (
	"testing"
)

func TestEncryptMessage(t *testing.T) {
	encryptedMessage := encryptMessage("Hello World", "Test passphrase!")
	plainTextMessage := decryptMessage(encryptedMessage, "Test passphrase!")
	if encryptedMessage != plainTextMessage {
		t.Errorf("Message was not successfully encrpyted")
	}
}
