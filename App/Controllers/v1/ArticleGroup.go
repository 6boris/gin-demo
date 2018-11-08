package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"net/http"
)

//	获取分组信息
func GetAllArticleGroup(c *gin.Context) {
	groups := Services.GetAllArticleGroup()

	c.JSON(http.StatusOK, gin.H{
		"name": "GetAllArticleGroup",
		"data": groups,
	})

}
