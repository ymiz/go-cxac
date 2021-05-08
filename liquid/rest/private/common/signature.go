package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateSignature(token *Token, path string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["path"] = path
	claims["nonce"] = time.Now().UnixNano()
	claims["token_id"] = token.Id()

	return t.SignedString([]byte(token.Secret()))
}
