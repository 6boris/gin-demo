package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllArticleGroup(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "GetAllArticleGroup",
	})
}
