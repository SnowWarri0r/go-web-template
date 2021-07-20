package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-web-template/config"
)

func Redis(ctx *gin.Context) *redis.Conn {
	return config.Redis(ctx)
}
