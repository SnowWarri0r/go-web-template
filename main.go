package main

import (
	"go-web-template/config"
	"go-web-template/router"
	"log"
)

func main() {
	config.InitConfig()
	config.LoadConfig()
	e := router.InitRouter()
	srv := config.InitServer(e)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n\n", err)
		}
	}()
	router.CloseRouter(srv)
}
