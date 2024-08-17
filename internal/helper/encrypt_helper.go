package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type OpenSslCipherSettings struct {
	CipherAlgorithm string
	Option          int
	IV              []byte
}

func getOpenSslCipherSettings() OpenSslCipherSettings {
	cipherAlgorithm := "aes-128-cbc"
	option := 0
	iv := []byte{0xe8, 0xe9, 0xed, 0xd2, 0xef, 0x02, 0xeb, 0xa7, 0x01} // Assuming the PHP byte string is in hexadecimal

	return OpenSslCipherSettings{
		CipherAlgorithm: cipherAlgorithm,
		Option:          option,
		IV:              iv,
	}
}

func Encrypt(str string, appKey string) (string, error) {
	settings := getOpenSslCipherSettings()
	key := []byte(appKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a new cipher block
	cipherText := make([]byte, len(str))
	iv := settings.IV
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, []byte(str))

	return fmt.Sprintf("%x", cipherText), nil
}

func Decrypt(cipherText string, appKey string) (string, error) {
	settings := getOpenSslCipherSettings()
	key := []byte(appKey)

	// Convert hex string to byte slice
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a new cipher block
	if len(cipherBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := settings.IV
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherBytes, cipherBytes)

	return string(cipherBytes), nil
}

func Base64Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Base64Decode(value string) (string, error) {
	buffer, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
