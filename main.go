package main

import (
	"Wizz-homepage-go/Global"
	"Wizz-homepage-go/models"
	"Wizz-homepage-go/route"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
)

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

	db := viper.Get("database")
	fmt.Printf("Using %v \n",db)

	if db == "sqlite3" {
		Global.Database, err = gorm.Open("sqlite3", "./wizz-homepage-backend.Database")
	} else if db == "mysql" {
		Global.Database, err = gorm.Open("mysql", getMysqlConnectString())
	}
	if err != nil {
		log.Println(err)
		log.Fatal("Database connect error")
	}
	defer Global.Database.Close()

	Global.Database.AutoMigrate(&models.Story{})
	Global.Database.AutoMigrate(&models.Product{})
	Global.Database.AutoMigrate(&models.Member{})

	route.ProcessRoute()

	Global.Engine.StaticFile("", "./html")
	Global.Engine.Static("static", "./html/static")

	_ = Global.Engine.Run()

}
