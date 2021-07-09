package router

import (
	"github.com/gin-gonic/gin"
	"go-web-template/controller"
)

func initTest(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/hello",controller.Hello())
	routerGroup.GET("/token",controller.Token())
}
