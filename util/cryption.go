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

// Encrypt encrypts the given plaintext using the provided key.
// The key should be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func Encrypt(plaintext string) (string, error) {
	key := global.SysConfig.CryPtKey
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext using the provided key.
// The key should be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func Decrypt(encrypted string) (string, error) {
	key := global.SysConfig.CryPtKey
	ciphertext, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
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
