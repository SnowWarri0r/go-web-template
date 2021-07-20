package config

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log1 "log"
)

type logg struct {
	log *log1.Logger
}

func (l *logg) Printf(ctx context.Context, format string, v ...interface{}) {
	_ = l.log.Output(2, fmt.Sprintf(format, v...))
}

var rdb *redis.Client

func InitRedis() {
	redis.SetLogger(&logg{log: log1.New(log.Writer(), "redis: ", log1.LstdFlags|log1.Lshortfile)})
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Conf.RedisUrl,
		Password: redisPasswd,
		DB:       0,
		PoolSize: config.Conf.RedisPoolSize,
	})
}

func Redis(ctx *gin.Context) *redis.Conn {
	return rdb.Conn(ctx)
}
func StopRedis() {
	if err := rdb.Close(); err != nil {
		Log().Warnln(err.Error())
	}
}
