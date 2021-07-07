package util

import (
	"testing"

	"github.com/kataras/jwt"
)

var (
	goodToken           []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwic3RhZ2UiOiIxIiwidXVpZCI6IjY0ZDM5YWIwLWRhNjUtMTFlYi04ODhmLTAyMDAxMzAyYmI0OCIsImlhdCI6MTYyNTE0MTQ1OCwiZXhwIjoxOTQxMDgxMjYxfQ.vWnJ2gMYt8hJnowhYr3LsyknWVLc5XuP0RWfjBFm870")
	wrongSignatureToken []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwic3RhZ2UiOiIxIiwidXVpZCI6IjY0ZDM5YWIwLWRhNjUtMTFlYi04ODhmLTAyMDAxMzAyYmI0OCIsImlhdCI6MTYyNTE0MTQ1OCwiZXhwIjoxOTQxMDgxMjYxfQ.vWnJ2gMYt8hJnowhYr3LsyknWVLc5XuP0RWfjBFn870")
	noUserIDToken       []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGFnZSI6IjEiLCJ1dWlkIjoiNjRkMzlhYjAtZGE2NS0xMWViLTg4OGYtMDIwMDEzMDJiYjQ4IiwiaWF0IjoxNjI1MTQxNDU4LCJleHAiOjE5NDEwODEyNjF9.AmOXFiAYrGICu3fj1Yt4dzxehWhAeuztqELO8cVGbg8")
	noStageToken        []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwidXVpZCI6IjY0ZDM5YWIwLWRhNjUtMTFlYi04ODhmLTAyMDAxMzAyYmI0OCIsImlhdCI6MTYyNTE0MTQ1OCwiZXhwIjoxOTQxMDgxMjYxfQ.59ZMk4TMjhRMcf7Ct-A5i3MS3dRI8kkWWPp0D4L0mGg")
	noUuidToken         []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwic3RhZ2UiOiIxIiwidXVpZCI6IiIsImlhdCI6MTYyNTE0MTQ1OCwiZXhwIjoxOTQxMDgxMjYxfQ.h2KnrjQnSM41eTo-2NOG1C3XrrNKi0AIY2dZPum2Wdk")
	expiredToken        []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwic3RhZ2UiOiIxIiwidXVpZCI6IjY0ZDM5YWIwLWRhNjUtMTFlYi04ODhmLTAyMDAxMzAyYmI0OCIsImlhdCI6MTYyNTE0MTQ1MCwiZXhwIjoxNjI1MTQxNDU4fQ.Kr1aA8d5_dUJ-nSve5f-oZVkT490CqON4xl9I5bQbTM")
)

func TestActiveTokens(t *testing.T) {
	t.Run("added token can be retrieved", func(t *testing.T) {
		jwtActiveTokens = make([][]byte, 0)

		t1, _ := GenJWTToken("1", "1")
		t2, _ := GenJWTToken("2", "2")
		t3, _ := GenJWTToken("3", "3")

		AddActiveToken(t1)
		AddActiveToken(t2)
		AddActiveToken(t3)

		if len(jwtActiveTokens) != 3 {
			t.Fatal("tokens must be active")
		}
	})
	t.Run("token cannot be added more than once", func(t *testing.T) {
		jwtActiveTokens = make([][]byte, 0)
		t1, _ := GenJWTToken("1", "1")

		AddActiveToken(t1)
		AddActiveToken(t1)

		if len(jwtActiveTokens) != 1 {
			t.Fatal("jwtActiveTokens must not contain duplicate tokens")
		}
	})
	t.Run("tokens can be removed", func(t *testing.T) {
		jwtActiveTokens = make([][]byte, 0)
		t1, _ := GenJWTToken("1", "1")

		AddActiveToken(t1)
		RemoveActiveToken(t1)

		if len(jwtActiveTokens) != 0 {
			t.Fatal("token must be removed from jwtActiveTokens")
		}
	})
}

func TestJWTToken(t *testing.T) {
	t.Run("should return a valid token", func(t *testing.T) {
		// jwt token must be generated correctly
		token, err := GenJWTToken("1", "1")
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
		if _, ok := claims[userIdClaim]; !ok {
			t.Fatal("token must contain user ID claim")
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
	t.Run("test token with missing user ID", func(t *testing.T) {
		if IsGoodToken(noUserIDToken) {
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

func TestShouldReturnUserIDFromToken(t *testing.T) {
	want := "5"
	token, err := GenJWTToken(want, "3")
	if err != nil {
		t.Fatal("could not create token")
	}
	got, err := GetUserIDFromToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("want %v but got %v", want, got)
	}

}
