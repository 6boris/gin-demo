package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/App/Repositories/Services"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	admin, err := Services.AdminLogin(email, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":     "登录失败",
			"err_msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"data": admin,
	})

}

func AdminRegister(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	admin, err := Services.AdminRegister(email, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":     "注册失败",
			"err_msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册失败",
		"data": admin,
	})

}
