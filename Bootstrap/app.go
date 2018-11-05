package Bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/Routes"
)

func GetApp() *gin.Engine {
	app := gin.New()

	app.Use(gin.Logger())

	app.Use(gin.Recovery())

	app = Routes.InitRouter(app)

	return app
}
