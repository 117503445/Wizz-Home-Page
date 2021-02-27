package boot

import (
	"github.com/gogf/gf/os/gfile"
	_ "wizz-home-page/library/cron"
	_ "wizz-home-page/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	_ = gfile.Mkdir("./tmp/file")

	LogBindEs()

	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
