package main

import (
	"github.com/kylesliu/gin-demo/Bootstrap"
	"net/http"
	"time"
)

func main() {
	router := Bootstrap.GetApp()

	app := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.ListenAndServe()
}
