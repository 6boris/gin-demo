package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/app/model"
	"net/http"
)

func IndexArticle(c *gin.Context) {
	article := []model.Article{}
	db := model.GetDB()
	db.Preload("Group").Find(&article)

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

func ShowArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "show",
	})
}

func EditArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "edit",
	})
}

func DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "delete",
	})
}
