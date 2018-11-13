package Routes

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Controllers/v1"
	"github.com/kylesliu/gin-demo/App/Controllers/v3"
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"github.com/kylesliu/gin-demo/Bootstrap/config"
	"log"
	"time"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(config.JwtConfig.Secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "a",

		//	自定义登录载体
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*MySQL.Admin); ok {
				return jwt.MapClaims{
					"Admin.IsLogin": true,
					"Admin.Id":      v.Id,
					"Admin.Name":    v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			return nil
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			email := c.Query("email")
			password := c.Query("password")

			data := c.PostFormMap("email")

			fmt.Println("asd", email, password, data)

			if admin, err := Services.AdminLoginCheck(email, password); err != nil {
				return nil, jwt.ErrFailedAuthentication
			} else {
				return &MySQL.Admin{
					Id:    admin.(*MySQL.Admin).Id,
					Email: admin.(*MySQL.Admin).Email,
					Name:  admin.(*MySQL.Admin).Name,
				}, nil
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	app.GET("/v1/articles", v1.GetAllArticle)
	app.GET("/v1/articles/:id", v1.GetOneArticle)
	app.GET("/v1/article_groups", v1.GetAllArticleGroup)
	app.POST("/v1/index/auth/login", v1.IndexLogin)

	app.POST("/v3/admin/user_info", v3.GetAllAdminInfo)
	app.POST("/v3/admin/auth/login", v3.AdminLogin)
	app.POST("/v3/admin/auth/register", v3.AdminRegister)

	apiv3 := app.Group("/api/v3")
	apiv3.POST("/admin/auth/login", authMiddleware.LoginHandler)
	apiv3.POST("/admin/auth/refresh_token", authMiddleware.RefreshHandler)
	apiv3.POST("/admin/auth/register", v3.AdminRegister)
	apiv3.Use(authMiddleware.MiddlewareFunc())
	{
		apiv3.GET("/admin/users", v3.GetAllAdminInfo)
	}
	return app
}
