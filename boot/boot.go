package boot

import (
	"wizz-home-page/app/model"
	_ "wizz-home-page/library/cron"
	_ "wizz-home-page/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	str, _ := model.EncryptPassword("admin")
	g.Log().Line().Debug(str)
	//middleware.NeedRole("123")
	LogBindEs()

	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
