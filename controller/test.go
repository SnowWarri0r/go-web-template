package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-template/service"
)

func Hello() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		result:=service.Hello()
		ok(ctx,result)
	}
}
