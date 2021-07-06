package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"reflect"
)

// key is used by the function inside this package
var key []byte = []byte("026BD6BFEE7B6D8128A1952E05A17E24")

// PrepareKey must assign a value to key variable (from a file or from an environment variable)
func PrepareKey() error {
	// TODO: implement method

	return errors.New("unimplemented method")
}

// EncryptData encrypts received data
func EncryptData(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciferData := aesGCM.Seal(nonce, nonce, data, nil)

	return ciferData, nil
}

// DecryptData decrypts received data using package key variable
func DecryptData(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}
	return plaintext, nil
}

// VerifyEncryptedData verify equality between received encrypted data and original data
func VerifyEncryptedData(encryptedData []byte, originalData []byte) bool {
	data, _ := DecryptData(encryptedData)
	return reflect.DeepEqual(data, originalData)
}
