package main

import (
	"github.com/kylesliu/gin-demo/Bootstrap"
)

func main() {
	router := Bootstrap.GetApp()
	router.ListenAndServe()
}
