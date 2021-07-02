package jwt

import (
	"testing"

	"github.com/kataras/jwt"
)

var (
	goodToken           []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6IjEiLCJ1dWlkIjoiNjRkMzlhYjAtZGE2NS0xMWViLTg4OGYtMDIwMDEzMDJiYjQ4IiwiaWF0IjoxNjI1MTQxNDU4LCJleHAiOjE2MjUyMjc4NTh9.LmnMc8Pa384-FTPnC3-OZVHeQBYVMUgIuZbpKVn5lMA")
	wrongSignatureToken []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6IjEiLCJ1dWlkIjoiNjRkMzlhYjAtZGE2NS0xMWViLTg4OGYtMDIwMDEzMDJiYjQ4IiwiaWF0IjoxNjI1MTQxNDU4LCJleHAiOjE2MjUyMjc4NTh9.LmnMc8Pa384-FTPnC3-OZVHeQBYVMUgIuZbpKVm5lMA")
	noStageToken        []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiNjRkMzlhYjAtZGE2NS0xMWViLTg4OGYtMDIwMDEzMDJiYjQ4IiwiaWF0IjoxNjI1MTQxNDU4LCJleHAiOjE2MjUyMjc4NTh9.RUt8P_Gptcio4Q-6tOv0kG-iFgC2JJ9npFPvkMdrqEQ")
	noUuidToken         []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6IjEiLCJpYXQiOjE2MjUxNDE0NTgsImV4cCI6MTYyNTIyNzg1OH0.H6UGJAhWgay1UT_0dj1a6VR4mx5sktZOdQ8vcnXgfJE")
	expiredToken        []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6IjEiLCJ1dWlkIjoiNjRkMzlhYjAtZGE2NS0xMWViLTg4OGYtMDIwMDEzMDJiYjQ4IiwiaWF0IjoxNjI1MTQxNDU4LCJleHAiOjE2MjUxNTg4MDB9.Hmo97VKMDENRVvCxA3nihpAePiqSjR2ujm1RrQC44CU")
)

func TestJWTToken(t *testing.T) {
	t.Run("should return a valid token", func(t *testing.T) {
		// jwt token must be generated correctly
		token, err := GenJWTToken("1")
		if err != nil {
			t.Fatal(err)
		}
		// jwt token must be valid
		verifiedToken, err := jwt.Verify(jwtAlg, jwtKey, token)
		if err != nil {
			t.Fatal(err)
		}
		// jwt token must contain stage and uuid
		claims := make(map[string]string)
		verifiedToken.Claims(&claims)
		if len(claims) == 0 {
			t.Fatal("invalid claims")
		}
		if _, ok := claims[stageClaim]; !ok {
			t.Fatal("token must contain stage claim")
		}
		if _, ok := claims[uuidClaim]; !ok {
			t.Fatal("token must contain uuid")
		}
	})
}

func TestIsGoodToken(t *testing.T) {
	t.Run("test a valid and complete token", func(t *testing.T) {
		if !IsGoodToken(goodToken) {
			t.Fatal("this token must be valid")
		}
	})
	t.Run("test token with wrong signature", func(t *testing.T) {
		if IsGoodToken(wrongSignatureToken) {
			t.Fatal("this token must be invalid")
		}
	})
	t.Run("test token with missing stage", func(t *testing.T) {
		if IsGoodToken(noStageToken) {
			t.Fatal("this token must be invalid")
		}
	})
	t.Run("test token with missing uuid", func(t *testing.T) {
		if IsGoodToken(noUuidToken) {
			t.Fatal("this token must be invalid")
		}
	})
	t.Run("test expired token", func(t *testing.T) {
		if IsGoodToken(expiredToken) {
			t.Fatal("expired token")
		}
	})
}
