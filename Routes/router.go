package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Controllers/v1"
	"github.com/kylesliu/gin-demo/App/Controllers/v3"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	app.GET("/v1/articles", v1.GetAllArticle)
	app.GET("/v1/articles/:id", v1.GetOneArticle)
	app.GET("/v1/article_groups", v1.GetAllArticleGroup)
	app.POST("/v1/index/auth/login", v1.IndexLogin)

	app.POST("/v3/admin/user_info", v3.GetAllAdminInfo)
	app.POST("/v3/admin/auth/login", v3.AdminLogin)
	app.POST("/v3/admin/auth/register", v3.AdminRegister)

	return app
}
