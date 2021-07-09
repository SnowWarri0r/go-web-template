package service

import (
	"github.com/dgrijalva/jwt-go"
	"go-web-template/config"
	"go-web-template/util"
	"time"
)

func Hello() util.Result {
	return util.SuccessInfo("Hello,World!")
}

func Token() util.Result {
	j := config.InitJWT()
	token, err := j.CreateToken(util.Claims{
		ID: 1,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "SnowWarrior",
			IssuedAt:  time.Now().Unix(),
		},
	})
	if err != nil {
		return util.FailureInfo("Creating token failed")
	}
	return util.SuccessData(token)
}
