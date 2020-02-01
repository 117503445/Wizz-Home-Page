package main

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Middlewares"
	"Wizz-Home-Page/models"
	"Wizz-Home-Page/route"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
)

//func getMysqlConnectString() string {
//	hostname := viper.Get("mysql.hostname")
//	port := viper.Get("mysql.port")
//	un := viper.Get("mysql.username")
//	pd := viper.Get("mysql.password")
//	dbName := viper.Get("mysql.databaseName")
//	connectString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", un, pd, hostname, port, dbName)
//	fmt.Print("mysql connect string is -->")
//	fmt.Println(connectString)
//	return connectString
//}

//判断 文件or路径 是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//如果path文件夹不存在才创建
func SafeMkdir(path string) {
	var err error
	fmt.Println(!Exists(path))
	if !Exists(path) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

// @title Wizz-Home-Page API
// @version 1.0
// @description Wizz's HomePage Backend

// @contact.name 117503445
// @contact.url https://github.com/117503445
// @contact.email t117503445@gmail.com

// @license.name GNU General Public License v3.0
// @license.url https://github.com/TGclub/Wizz-Home-Page/blob/master/LICENSE

// @host localhost:8080
// @BasePath /api
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	var err error

	SafeMkdir("./data")
	SafeMkdir("./data/logs")

	//向 ./data/logs/***.log 记录日志
	// *** 比如为 2020-01-31-13-34-55
	f, err := os.OpenFile("./data/logs/"+time.Now().Format("2006-01-02-15-04-05")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	} else {
		log.Println(err)
	}

	//从 config.json 读取配置
	viper.SetConfigFile("config.json")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config error", err)
	}

	Middlewares.NameAndPassword = viper.GetStringMapString("account")

	//db := viper.Get("database")
	//fmt.Printf("Using %v \n", db)
	//
	//if db == "sqlite3" {
	//	Global.Database, err = gorm.Open("sqlite3", "./data/Wizz-Home-Page.Database")
	//} else if db == "mysql" {
	//	Global.Database, err = gorm.Open("mysql", getMysqlConnectString())
	//}
	Global.Database, err = gorm.Open("sqlite3", "./data/Wizz-Home-Page.Database")
	if err != nil {
		log.Fatal(err)
	}

	Global.Database.AutoMigrate(&models.Story{})
	Global.Database.AutoMigrate(&models.Product{})
	Global.Database.AutoMigrate(&models.Member{})
	Global.Database.AutoMigrate(&models.ServerLog{})

	Global.Engine = gin.Default()

	route.InitRoute()

	//接入前端
	Global.Engine.StaticFile("", "./html")
	Global.Engine.Static("static", "./html/static")

	err = Global.Engine.Run()
	if err != nil {
		log.Println(err)
	}
	err = Global.Database.Close()
	if err != nil {
		panic(err)
	}
}