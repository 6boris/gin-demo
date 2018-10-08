package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	SecretKey = "welcome to wangshubo's blog"
)

type Token struct {
	Token string `json:"token"`
}

func JwtLogin(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = "kyle"
	claims["aud"] = "test"
	claims["admin"] = true
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		fmt.Println(err)
	}
	//response := Token{tokenString}

	c.JSON(http.StatusOK, gin.H{
		"data": Token{tokenString},
	})
}

func JwtCheckLogin(c gin.Context) {

}
