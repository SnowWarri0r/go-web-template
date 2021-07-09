package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer(handler *gin.Engine) *http.Server{
	return &http.Server{
		Addr: Addr,
		Handler: handler,
	}
}
