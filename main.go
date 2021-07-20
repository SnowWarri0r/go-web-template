package main

import (
	"go-web-template/config"
	"go-web-template/router"
)

func main() {
	config.InitConfig()
	config.LoadConfig()
	config.InitLogger()
	config.InitDB()
	config.InitRedis()
	e := router.InitRouter()
	srv := config.InitServer(e)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			config.Log().Fatalf("listen: %s\n\n", err)
		}
	}()
	router.CloseRouter(srv)
	config.StopRedis()
	config.DisconnectDB()
}
