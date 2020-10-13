package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// AESEncrypterWithCFB encrypt a plain string into cipher with AES CFB mode
func AESEncrypterWithCFB(plainText string, key []byte) (cipherText string, err error) {
	plainBytes := []byte(plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherBytes := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], plainBytes)
	cipherText = hex.EncodeToString(cipherBytes)
	return cipherText, nil
}

// AESDecrypterWithCFB decrypt a cipher string to plain with AES CFB mode
func AESDecrypterWithCFB(cipherText string, key []byte) (plainText string, err error) {
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherBytes) < aes.BlockSize {
		return "", errors.New("cipher text is too short")
	}
	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)
	plainText = string(cipherBytes)
	return plainText, nil
}

// ParseKey convert a string into a valid key bytes
func ParseKey(key string) (keyBytes []byte, err error) {
	str := base64.StdEncoding.EncodeToString([]byte(key))
	keyBytes = []byte(str)
	if len(keyBytes) < 16 {
		return nil, aes.KeySizeError(len(keyBytes))
	}
	if len(keyBytes) < 24 {
		return keyBytes[:16], nil
	}
	if len(keyBytes) < 32 {
		return keyBytes[:24], nil
	}
	return keyBytes[:32], nil
}
