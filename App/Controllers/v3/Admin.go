package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"net/http"
)

func GetAllAdminInfo(c *gin.Context) {

	admins := Services.GetAdminInfo()

	c.JSON(http.StatusOK, gin.H{
		"name": "GetAllAdminInfo",
		"data": admins,
	})

}
