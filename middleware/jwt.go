package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/config"
	"go-web-template/controller"
	"go-web-template/util"
	"strings"
)

func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authString := ctx.GetHeader("Authorization")
		if len(authString) < 1 {
			controller.Forbidden(ctx)
			ctx.Abort()
		}
		tokenArray := strings.Split(authString, " ")
		if len(tokenArray[1]) < 1 {
			controller.Unauthorized(ctx)
			ctx.Abort()
		}
		token := tokenArray[1]
		j := config.InitJWT()
		claim, err := j.ParserToken(token)
		if err != nil {
			switch err {
			case util.TokenExpired, util.TokenNotValidYet, util.TokenInvalid, util.TokenMalformed:
				controller.Unauthorized(ctx)
				ctx.Abort()
			}
		}
		ctx.Set("claim", claim)
		ctx.Next()
	}
}
