package config

import (
	"github.com/gin-gonic/gin"
	"go-web-template/util"
	"net/http"
)

func InitServer(handler *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func InitJWT() *util.JWT {
	return &util.JWT{
		Salt: []byte(salt),
	}
}
