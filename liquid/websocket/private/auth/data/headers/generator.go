package headers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ymiz/go-cxac/liquid/rest/private/common"
	"time"
)

func GenerateHeaders(token *common.Token, path string) (*Headers, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["path"] = path
	claims["token_id"] = token.Id()
	claims["nonce"] = time.Now().UnixNano()

	sign, err := t.SignedString([]byte(token.Secret()))
	if err != nil {
		return nil, err
	}
	return NewHeaders(sign), nil
}
