package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/luolayo/gin-study/global"
	"io"
)

var key = "=CwuV1AsX_nI6KN2iX0NT2XxY-m-Itx8"

// Encrypt encrypts the given plaintext using the provided key.
// The key should be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		global.LOG.Error("Failed to create cipher: %v", err)
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		global.LOG.Error("Failed to read random: %v", err)
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext using the provided key.
// The key should be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func Decrypt(encrypted string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		global.LOG.Error("Failed to decode base64: %v", err)
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		global.LOG.Error("Failed to create cipher: %v", err)
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		global.LOG.Error("Ciphertext too short")
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// Compare compares the given plaintext with the given ciphertext.
func Compare(encryptedCharacter, unencryptedCharacter string) (bool, error) {
	decryptedCharacter, err := Decrypt(encryptedCharacter)
	if err != nil {
		return false, err
	}
	return decryptedCharacter == unencryptedCharacter, nil
}
