package util

import (
	"log"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPrepareKey(t *testing.T) {
	if err := PrepareKey(); err != nil {
		t.Fatal(err)
	}
}

func TestEncryptingAndDecryping(t *testing.T) {
	t.Run("test should encrypt data", func(t *testing.T) {
		oData := "3" // original data
		eData, err := EncryptData([]byte(oData))
		if err != nil {
			t.Fatal(err)
		}
		if eData == nil {
			t.Fatal("encrypted data could not be nil")
		}
		if oData == string(eData) {
			t.Fatal("original data was not encrypted")
		}
	})
	t.Run("test should decrypt encrypted data", func(t *testing.T) {
		want := []byte("3")
		eData := []byte{0x6d, 0x9e, 0xb3, 0xec, 0xa4, 0x12, 0x28, 0xe1, 0xae, 0xd8, 0x5d, 0xd0, 0x3f, 0xb5, 0x7a, 0x8e, 0x98, 0x13, 0xbd, 0xa9, 0xac, 0xc2, 0x5b, 0x42, 0xd5, 0x06, 0xea, 0xf2, 0x86}
		got, err := DecryptData(eData)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("want %v but got %v", want, got)
		}
	})
}

// TestHashing is only used to demonstrate a hashing concept
func TestHashing(t *testing.T) {
	t.Run("generate hash password", func(t *testing.T) {
		passwords := []string{"Popescu", "Zbagan", "Csismar", "Bitoanca", "Zbagan", "Zanfir", "Tehanciuc", "Siclovan"}
		for _, v := range passwords {
			log.Println(v)
			hashPwd, _ := bcrypt.GenerateFromPassword([]byte(v), 0)
			log.Println(string(hashPwd))
		}
	})
	t.Run("test should generate a hashed password", func(t *testing.T) {
		password := []byte("some very secured password")

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			t.Fatal(err)
		}
		if hashedPassword == nil {
			t.Fatal("hash password cannot be nil")
		}
	})
	t.Run("test unhashed password must be equal to original password", func(t *testing.T) {
		password := []byte("some very secure password")

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			t.Fatal(err)
		}
		if bcrypt.CompareHashAndPassword(hashedPassword, password) != nil {
			t.Fatal("unhashed password and original password do not match")
		}
	})
}
