package v1

import (
	"fmt"
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
	article_id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(c.Param("id"))

	article := Services.GetOneArticle(article_id)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功:GetOneArticle",
		"data": article,
	})
}
