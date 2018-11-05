package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"net/http"
)

func GetAllArticle(c *gin.Context) {
	//group_id := c.Query("group_id")

	articles := Services.GetAllArticle()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetAllArticle",
		"data": articles,
	})
}

func GetOneArticle(c *gin.Context) {
	article_id := c.Query("article_id")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetOneArticle",
		"data": article_id,
	})
}
