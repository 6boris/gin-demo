package Api

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"login": "login",
	})
}
