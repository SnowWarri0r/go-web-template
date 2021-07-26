package dao

import (
	"github.com/gin-gonic/gin"
	"go-web-template/config"
	"gorm.io/gorm"
)

func DB(ctx *gin.Context) *gorm.DB {
	return config.DB(ctx)
}
