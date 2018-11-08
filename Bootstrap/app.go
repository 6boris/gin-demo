package Bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"github.com/kylesliu/gin-demo/Bootstrap/config"
	"github.com/kylesliu/gin-demo/Routes"
	"net/http"
	"strconv"
	"time"
)

func GetApp() *http.Server {
	route := gin.New()
	//gin.SetMode(config.ServerSetting.RunMode)
	gin.SetMode("debug")
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	config.Setup()
	Services.MySQLSetup()

	route = Routes.InitRouter(route)

	server := &http.Server{
		Addr:           config.ServerConfig.HttpAddr + ":" + strconv.Itoa(config.ServerConfig.HttpPort),
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}
