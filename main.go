package main

import (
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

var db *gorm.DB


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
func getMysqlConnectString()string{
	hostname:=viper.Get("mysql.hostname")
	port:=viper.Get("mysql.port")
	un:=viper.Get("mysql.username")
	pd:=viper.Get("mysql.password")
	dbName:=viper.Get("mysql.databaseName")
	connectString:=fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",un,pd,hostname,port,dbName)
	return connectString
}
func main() {
	var err error

	viper.SetConfigFile("config.json")

	err=viper.ReadInConfig()
	if err!=nil{
		log.Println("read config error",err)
	}

	fmt.Println(viper.Get("mysql.username"))
	
	//db, err = gorm.Open("sqlite3", "./wizz-homepage-backend.db")
	db, err = gorm.Open("mysql", getMysqlConnectString())
	//todo: edit mysql string
	if err != nil {
		log.Println(err)
		log.Fatal("Database connect error")
	}
	defer db.Close()
	db.AutoMigrate(&models.Story{})

	//create the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"), //todo:Edit secret key
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour * 24 * 7,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	engine := gin.Default()

	apiGroup := engine.Group("/api")

	apiGroup.GET("/stories", readStories)
	apiGroup.GET("/stories/:id", readStory)
	apiGroup.POST("/stories", authMiddleware.MiddlewareFunc(), createStory)
	apiGroup.PUT("/stories/:id", authMiddleware.MiddlewareFunc(), updateStory)
	apiGroup.DELETE("/stories/:id", authMiddleware.MiddlewareFunc(), deleteStory)

	auth := apiGroup.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)

	_ = engine.Run()

}

func readStories(c *gin.Context) {
	var stories []models.Story
	db.Find(&stories)
	c.JSON(200, stories)
}

func readStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	db.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	}
	c.JSON(200, story)
}

func createStory(c *gin.Context) {
	var story models.Story
	_ = c.BindJSON(&story) //绑定一个请求主体到一个类型
	db.Create(&story)
	c.JSON(200, story)
}

func updateStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	db.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	} else {
		_ = c.BindJSON(&story)
		db.Save(&story)
		c.JSON(200, story)
	}
}

func deleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	db.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	} else {
		_ = c.BindJSON(&story)
		db.Delete(&story)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}