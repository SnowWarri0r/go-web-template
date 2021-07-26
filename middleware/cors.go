package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("origin"))
			ctx.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Access-Control-Request-Headers,Access-Control-Request-Method,Origin,Referer,Sec-Fetch-Dest,Accept-Language,Accept-Encoding,Sec-Fetch-Mode,Sec-Fetch-Site,User-Agent")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			ctx.Header("Access-Control-Max-Age", "1728000")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
