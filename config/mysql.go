package config

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func InitDB() {
	var err error
	var logLevel logger.LogLevel
	var sqlDB *sql.DB
	if mode == "prod" {
		logLevel = logger.Silent
	} else {
		logLevel = logger.Info
	}
	newLogger := logger.New(Log(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		})
	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //建表命名约束为根据结构体名称单数小写命名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: mysqlUser + ":" + mysqlPasswd + "@tcp(" + dbUrl + ")/" + mysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &mysqlConfig)
	if err != nil {
		Log().Fatalln(err.Error())
		return
	}
	sqlDB, err = db.DB()
	if err != nil {
		Log().Fatalln(err.Error())
		return
	}
	//设置连接池最大连接数
	sqlDB.SetMaxOpenConns(mysqlMaxOpenConns)
	//设置连接池最大闲置连接数
	sqlDB.SetMaxIdleConns(mysqlMaxIdleConns)
	//设置连接最大闲置时间
	sqlDB.SetConnMaxIdleTime(time.Duration(mysqlConnMaxIdleTime))
	//设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConnMaxLifeTime))
	// 自动迁移设置todo
}

func DB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}
// DisconnectDB 断开与数据库的连接
func DisconnectDB() {
	mysqlDB, _ := db.DB()
	if err := mysqlDB.Close(); err != nil {
		Log().Warnln(err.Error())
	}
}