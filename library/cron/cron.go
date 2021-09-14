package cron

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"wizz-home-page/app/service/serverchan"
	"wizz-home-page/app/service/wjx"
)

func init() {
	g.Log().Line().Debug("cron load")
	// 0 * * * * * every minutes
	if _, err := gcron.Add("@every 60s", func() { wjx.PullResume() }, "wjx-PullResume"); err != nil {
		g.Log().Line().Error(err)
	}
	// 0 0 21 * * * 每晚9点执行此操作
	if _, err := gcron.Add("0 0 21 * * *", func() { serverchan.ResumeRemind() }, "serverchan-ResumeRemind"); err != nil {
		g.Log().Line().Error(err)
	}
}
