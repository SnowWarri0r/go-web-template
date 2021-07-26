package config

import (
	"github.com/gin-gonic/gin"
	"go-web-template/util"
	log2 "log"
	"net/http"
)

func InitServer(handler *gin.Engine) *http.Server {
	return &http.Server{
		Addr:     config.Server.Addr,
		ErrorLog: log2.New(log.Writer(), "", 0),
		Handler:  handler,
	}
}

func InitJWT() *util.JWT {
	return &util.JWT{
		Salt: []byte(config.Server.Salt),
	}
}
