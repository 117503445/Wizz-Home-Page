package boot

import (
	_ "wizz-home-page/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	LogBindEs()

	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
