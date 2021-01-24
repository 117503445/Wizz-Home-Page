package main

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Init"
	"github.com/spf13/viper"
	"log"
)

// @title Wizz-Home-Page API
// @version 1.2.1
// @description Wizz's HomePage Backend

// @contact.name 117503445
// @contact.url https://github.com/117503445
// @contact.email t117503445@gmail.com

// @license.name GNU General Public License v3.0
// @license.url https://github.com/TGclub/Wizz-Home-Page/blob/master/LICENSE

// @host homepage.backend.wizzstudio.com
// @BasePath /api
// @schemes https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	Init.Init()
	useHttps := viper.GetBool("useHttps")
	var err error
	if useHttps {
		err = Global.Engine.RunTLS(":443", "ssl.pem", "ssl.key")
	} else {
		err = Global.Engine.Run(":80")
	}
	if err != nil {
		log.Println(err)
	}
	err = Global.Database.Close()
	if err != nil {
		panic(err)
	}
}
