package Middlewares

import (
	"Wizz-Home-Page/Global"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type login struct {
	Username string `example:"admin" json:"username"`
	Password string `example:"admin" json:"password"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func getRandomBytes(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

func getSecretKey() []byte {
	pathKey := "./data/key.txt"
	b, err := ioutil.ReadFile(pathKey)
	if err == nil {
		return b
	} else {
		b := getRandomBytes(256)
		err = ioutil.WriteFile(pathKey, b, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		return b
	}
}

func payloadFunc(data interface{}) jwt.MapClaims {
	fmt.Println("PayloadFunc")
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}
func identityHandler(c *gin.Context) interface{} {
	fmt.Println("IdentityHandler")
	claims := jwt.ExtractClaims(c)
	return &User{
		UserName: claims[identityKey].(string),
	}
}

// @Summary 登录
// @Description 更改请求中的 Username 和 Password 进行登录。登陆成功以后，返回json中token字段比如说是"token":"123"，就在右上角Authorize按钮点一下，输入Bearer 123，大小写、空格敏感。然后就能使用需要身份验证的接口啦。
// @Tags 身份验证
// @Accept  json
// @Produce  json
// @Param   loginVals      body route.login true  "登录值"
// @Success 200 {string} string "{"code":200,"expire":"2020-02-05T23:11:41+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5MTU1MDEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDMxMDcwMX0.GWlmyTfCkXQYwgbtuTgVSTUSJXDcoDb_bptgRpt4HCU"}"
// @Router /auth/login [POST]
func authenticator(c *gin.Context) (interface{}, error) {
	fmt.Println("Authenticator")
	var loginVal login
	if err := c.ShouldBindJSON(&loginVal); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVal.Username
	password := loginVal.Password

	//if userID == "admin" && password == "admin" {
	//	return &User{
	//		UserName:  userID,
	//		LastName:  "Bo-Yi",
	//		FirstName: "Wu",
	//	}, nil
	//}

	if Global.NameAndPassword[userID] == password {
		return &User{ //todo 修改User
			UserName:  userID,
			LastName:  "Bo-Yi",
			FirstName: "Wu",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
func authorizator(data interface{}, c *gin.Context) bool {
	fmt.Print("Authorizator->")
	fmt.Println(data)
	if v, ok := data.(*User); ok && v.UserName == "admin" {
		return true
	}
	//todo:优化语句结构
	return false
}
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "wizz zone",
		Key:             getSecretKey(),
		Timeout:         time.Hour * 24 * 7,
		MaxRefresh:      time.Hour * 24 * 7,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		panic(err)
	}
	return authMiddleware
}
