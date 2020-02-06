package Init

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Middlewares"
	"Wizz-Home-Page/apis"
	"Wizz-Home-Page/models"
	"Wizz-Home-Page/route"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
)

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
	if !Exists(path) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
func Init() {

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
	viper.SetConfigFile("./config.json")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config error", err)
	}

	Middlewares.NameAndPassword = viper.GetStringMapString("account")

	apis.Ak = viper.GetString("qiniu.ak")
	apis.Sk = viper.GetString("qiniu.sk")
	apis.Bucket = viper.GetString("qiniu.bucket")
	apis.Place=viper.GetString("qiniu.place")
	apis.Domain=viper.GetString("qiniu.domain")
	apis.BackGroundImageUrls = viper.GetStringSlice("backgroundImageUrls")



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
}
