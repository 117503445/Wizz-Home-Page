package route

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/apis"
	_ "Wizz-Home-Page/docs"//引入 swagger 必备
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"time"
)

type login struct {
	Username string `example:"admin" form:"username" json:"username" binding:"required"`
	Password string `example:"admin" form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
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
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   loginVals      body route.login true  "登录值" default({"password":"admin","username":"admin"})
// @Success 200 {string} string "{"code":200,"expire":"2020-02-05T23:11:41+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5MTU1MDEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDMxMDcwMX0.GWlmyTfCkXQYwgbtuTgVSTUSJXDcoDb_bptgRpt4HCU"}"
// @Router /auth/login [POST]
func authenticator(c *gin.Context) (interface{}, error) {
	fmt.Println("Authenticator")
	var loginVals login
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if userID == "admin" && password == "admin" { //todo:Edit password
		return &User{
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
func ProcessRoute() {
	//create the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"), //todo:Edit secret key
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
		log.Fatal("JWT Error:" + err.Error())
	}

	apiGroup := Global.Engine.Group("/api")

	auth := apiGroup.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)

	storyGroup := apiGroup.Group("/stories")
	storyGroup.GET("", apis.ReadStories)
	storyGroup.GET("/:id", apis.ReadStory)
	storyGroup.POST("", authMiddleware.MiddlewareFunc(), apis.CreateStory)
	storyGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), apis.UpdateStory)
	storyGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), apis.DeleteStory)

	productGroup := apiGroup.Group("/products")
	productGroup.GET("", apis.ReadProducts)
	productGroup.GET("/:id", apis.ReadProduct)
	productGroup.POST("", authMiddleware.MiddlewareFunc(), apis.CreateProduct)
	productGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), apis.UpdateProduct)
	productGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), apis.DeleteProduct)

	memberGroup := apiGroup.Group("/members")
	memberGroup.GET("", apis.ReadMembers)
	memberGroup.GET("/:id", apis.ReadMember)
	memberGroup.POST("", authMiddleware.MiddlewareFunc(), apis.CreateMember)
	memberGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), apis.UpdateMember)
	memberGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), apis.DeleteMember)

	Global.Engine.GET("/ver", func(c *gin.Context) {
		c.JSON(200, "0130-1921")
	})

	Global.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
