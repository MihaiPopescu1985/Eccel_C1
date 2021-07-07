package util

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/jwt"
)

const (
	userIdClaim = "userId"
	stageClaim  = "stage"
	uuidClaim   = "uuid"
)

var (
	jwtKey          []byte   = []byte("myverrysecretkey)")
	jwtAlg          jwt.Alg  = jwt.HS256
	jwtExp                   = time.Now().Add(24 * time.Hour).Unix()
	jwtActiveTokens [][]byte = make([][]byte, 0)
)

// GenJWTToken generate a signed jwt token
// token is valid for $jwtExp period and contains a uuid
func GenJWTToken(userId, stage string) ([]byte, error) {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	claim := map[string]string{userIdClaim: userId, stageClaim: stage, uuidClaim: uuid.String()}

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
	tkUserID, ok := claims[userIdClaim]
	if !ok || tkUserID == "" {
		log.Println("invalid user ID")
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

func GetUserIDFromToken(token []byte) (string, error) {
	claims := make(map[string]interface{})
	verifiedToken, err := jwt.Verify(jwtAlg, jwtKey, token)
	if err != nil {
		log.Println("invalid token")
		return "", errors.New("invalid token")
	}
	if err := verifiedToken.Claims(&claims); err != nil {
		log.Println("invalid claims in token")
		return "", errors.New("invalid claims")
	}
	if tkUserId, ok := claims[userIdClaim]; !ok {
		return "", errors.New("missing user ID from token")
	} else {
		return fmt.Sprint(tkUserId), nil
	}
}

func AddActiveToken(token []byte) {
	// verify if token already exist
	if !IsTokenActive(token) {
		// than add token
		jwtActiveTokens = append(jwtActiveTokens, token)
	}
}

func IsTokenActive(token []byte) bool {
	idFromToken, _ := GetUserIDFromToken(token)

	for _, t := range jwtActiveTokens {
		if reflect.DeepEqual(t, token) {
			return true
		}
		idFromT, _ := GetUserIDFromToken(t)
		if idFromToken == idFromT {
			return true
		}
	}
	return false
}

func RemoveActiveToken(token []byte) {
	idFromToken, _ := GetUserIDFromToken(token)

	for i, t := range jwtActiveTokens {
		idFromT, _ := GetUserIDFromToken(t)
		if idFromToken == idFromT {
			jwtActiveTokens[i] = jwtActiveTokens[len(jwtActiveTokens)-1]
			jwtActiveTokens = jwtActiveTokens[:len(jwtActiveTokens)-1]
		}
	}
}
