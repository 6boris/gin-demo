package Middlewares

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var authMiddleware *jwt.GinJWTMiddleware
var err error
func init() {
	// the jwt middleware
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func CheckJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware.MaxRefresh = time.Hour
		c.Next()
	}
}
