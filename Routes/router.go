package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Controllers/v1"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	app.GET("/v1/articles", v1.GetAllArticle)
	app.GET("/v1/articles/:id", v1.GetOneArticle)
	app.GET("/v1/article_groups", v1.GetAllArticleGroup)
	app.POST("/v1/index/auth/login", v1.IndexLogin)

	return app
}
