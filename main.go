package main

import (
	"Wizz-homepage-go/Global"
	story "Wizz-homepage-go/apis"
	"Wizz-homepage-go/models"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var identityKey = "id"

func getMysqlConnectString() string {
	hostname := viper.Get("mysql.hostname")
	port := viper.Get("mysql.port")
	un := viper.Get("mysql.username")
	pd := viper.Get("mysql.password")
	dbName := viper.Get("mysql.databaseName")
	connectString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", un, pd, hostname, port, dbName)
	fmt.Print("mysql connect string is -->")
	fmt.Println(connectString)
	return connectString
}
func main() {
	var err error

	viper.SetConfigFile("config.json")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("read config error", err)
	}

	//fmt.Println(viper.Get("mysql.username"))

	//Database, err = gorm.Open("sqlite3", "./wizz-homepage-backend.Database")
	Global.Database, err = gorm.Open("mysql", getMysqlConnectString())
	//todo: edit mysql string
	if err != nil {
		log.Println(err)
		log.Fatal("Database connect error")
	}
	defer Global.Database.Close()
	Global.Database.AutoMigrate(&models.Story{})

	//create the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"), //todo:Edit secret key
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour * 24 * 7,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("PayloadFunc")
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
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
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Print("Authorizator->")
			fmt.Println(data)
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			//todo:优化语句结构
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	engine := gin.Default()


	apiGroup := engine.Group("/api")

	apiGroup.GET("/stories", story.ReadStories)
	apiGroup.GET("/stories/:id", story.ReadStory)
	apiGroup.POST("/stories", authMiddleware.MiddlewareFunc(), story.CreateStory)
	apiGroup.PUT("/stories/:id", authMiddleware.MiddlewareFunc(), story.UpdateStory)
	apiGroup.DELETE("/stories/:id", authMiddleware.MiddlewareFunc(), story.DeleteStory)

	auth := apiGroup.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)

	//engine.StaticFile("","./html")

	_ = engine.Run()

}
