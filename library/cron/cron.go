package cron

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"wizz-home-page/app/service/wjx"
)

func init() {
	if _, err := gcron.Add("@hourly", func() { wjx.Download() }, "wjx_download"); err != nil {
		g.Log().Line().Error(err)
	}
}
