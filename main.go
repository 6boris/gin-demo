package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kylesliu/gin-demo/app/controller"
	"github.com/kylesliu/gin-demo/app/middleware"
)

func main() {
	router := gin.New()

	//	开启日志
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//	路由
	v1 := router.Group("/v1")
	{
		v1.GET("/index", controller.IndexArticle)
		v1.GET("/show", controller.ShowArticle)
		v1.POST("/edit", controller.EditArticle)
		v1.POST("/delete", controller.DeleteArticle)
	}

	v2 := router.Group("/v2")

	v2.Use(middleware.JWTAuthCheck())
	{
		v2.GET("/index", controller.IndexArticle)
		v2.GET("/show", controller.ShowArticle)
		v2.POST("/edit", controller.EditArticle)
		v2.POST("/delete", controller.DeleteArticle)

		v2.GET("/auth/login", controller.JwtLogin)

	}

	router.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080

}
