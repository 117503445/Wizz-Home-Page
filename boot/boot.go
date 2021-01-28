package boot

import (
	"wizz-home-page/app/middleware"
	_ "wizz-home-page/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	middleware.NeedRole("123")
	LogBindEs()

	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
