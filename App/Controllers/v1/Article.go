package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"net/http"
	"strconv"
)

func GetAllArticle(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Query("group_id"))
	articles := Services.GetAllArticle(group_id)

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
