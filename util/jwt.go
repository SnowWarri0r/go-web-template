package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("Token 已经过期")
	TokenNotValidYet = errors.New("Token 未激活")
	TokenMalformed   = errors.New("Token 错误")
	TokenInvalid     = errors.New("Token 无效")
)

type JWT struct {
	Salt []byte
}
type Claims struct {
	ID uint
	jwt.StandardClaims
}

func (j *JWT) CreateToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(j.Salt)
}
func (j *JWT) ParserToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Salt, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c, nil
	}
	return nil, TokenInvalid
}
