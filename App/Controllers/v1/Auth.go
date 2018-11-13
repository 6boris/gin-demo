package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin-demo/Bootstrap/config"
	"net/http"
	"time"
)

func IndexLogin(c *gin.Context) {

	mySigningKey := []byte(config.JwtConfig.Secret)

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			Issuer:    "admin",                   //	[iss]	jwt签发着
			Subject:   "user",                    //	[sub]	jwt面向用户
			Audience:  "1",                       //	[aud]	jwt接受者
			ExpiresAt: time.Now().Unix(),         //	[exp]	jwt的过期时间，这个过期时间必须要大于签发时间
			Id:        "index.auth",              //	[jti]	jwt唯一身份标识
			IssuedAt:  time.Now().Unix() + 10000, //	[iat]	jwt签发时间
			NotBefore: time.Now().Unix(),         //	[nbf]	定义在什么时间之前，该jwt都是不可用的.
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	c.JSON(http.StatusOK, gin.H{
		"name":  "IndexLogin",
		"token": token,
		"ss":    ss,
	})
}

func TakeJwtToken(c *gin.Context) {

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			Issuer:    "admin",                 //	[iss]	jwt签发着
			Subject:   "user",                  //	[sub]	jwt面向用户
			Audience:  "1",                     //	[aud]	jwt接受者
			ExpiresAt: time.Now().Unix() + 20, //	[exp]	jwt的过期时间，这个过期时间必须要大于签发时间
			Id:        "index.auth",            //	[jti]	jwt唯一身份标识
			IssuedAt:  time.Now().Unix(),       //	[iat]	jwt签发时间
			NotBefore: time.Now().Unix(),       //	[nbf]	定义在什么时间之前，该jwt都是不可用的.
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.JwtConfig.Secret))
	fmt.Printf("%v %v", ss, err)

	c.JSON(http.StatusOK, gin.H{
		"name":  "IndexLogin",
		"token": token,
		"ss":    ss,
	})
}

