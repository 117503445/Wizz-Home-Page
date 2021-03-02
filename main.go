package main

import (
	_ "wizz-home-page/boot"
	_ "wizz-home-page/router"

	"github.com/gogf/gf/frame/g"
)

// @title       wizz-home-page API
// @version     1.8.14
// @description `wizz-home-page` 企业官网

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	s := g.Server()
	s.SetDumpRouterMap(false)
	s.SetIndexFolder(true)
	s.SetServerRoot("./tmp/file")
	s.Run()
}
