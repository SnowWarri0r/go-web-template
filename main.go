package main

import (
	"go-web-template/router"
	"log"
	"net/http"
)

func main() {
	e := router.InitRouter()
	srv := &http.Server{
		Addr:    ":8000",
		Handler: e,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n\n", err)
		}
	}()
	router.CloseRouter(srv)
}
