package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Controllers/v1"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	app.GET("/v1/articles", v1.GetAllArticle)
	app.GET("/v1/article", v1.GetOneArticle)
	app.GET("/v1/articles/groups", v1.GetAllArticleGroup)

	return app
}