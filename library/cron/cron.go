package cron

import (
	"wizz-home-page/app/service/wjx"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
)

func init() {
	g.Log().Line().Debug("cron load")
	// 0 * * * * * every minutes
	if _, err := gcron.Add("@every 10s", func() { wjx.PullResume() }, "wjx-PullResume"); err != nil {
		g.Log().Line().Error(err)
	}
}
