package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/jwt"
)

const (
	stageClaim = "stage"
	uuidClaim  = "uuid"
)

var (
	jwtKey []byte  = []byte("myverrysecretkey)")
	jwtAlg jwt.Alg = jwt.HS256
	jwtExp         = time.Now().Add(24 * time.Hour).Unix()
)

// GenJWTToken generate a signed jwt token
// token is valid for 24 hours and contains a uuid
func GenJWTToken(stage string) ([]byte, error) {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	claim := map[string]string{stageClaim: stage, uuidClaim: uuid.String()}

	token, err := jwt.Sign(jwtAlg,
		jwtKey,
		claim,
		jwt.Claims{
			Expiry:   jwtExp,
			IssuedAt: time.Now().Unix(),
		})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// isGoodToken returns true if token was not modified and the claims exists
func IsGoodToken(token []byte) bool {

	tk, err := jwt.Verify(jwtAlg, jwtKey, token)
	if err != nil {
		log.Println(err)
		return false
	}

	claims := make(map[string]interface{})
	if err := tk.Claims(&claims); err != nil {
		log.Println(err)
		return false
	}
	tkStage, ok := claims[stageClaim]
	if !ok || tkStage == "" {
		log.Println("invalid stage claim")
		return false
	}
	tkUuid, ok := claims[uuidClaim]
	if _, err = uuid.Parse(fmt.Sprint(tkUuid)); err != nil {
		log.Println(err)
		return false
	}
	if !ok {
		log.Println("invalid uuid claim")
		return false
	}
	return true
}
