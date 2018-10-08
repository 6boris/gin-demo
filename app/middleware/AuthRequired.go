package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTAuthCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{
		//	"message": "未认证",
		//})

		//context.Abort()
	}

}
