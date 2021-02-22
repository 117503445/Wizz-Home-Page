package serverchan

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/test/gtest"
	"testing"
	"wizz-home-page/boot"
)

func TestPush(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := "chanid.txt"
		if !gfile.Exists(path) {
			g.Log().Line().Error(path + " don't exist")
			t.Fail()
		}
		chanId := gfile.GetContents(path) // chanId 为敏感数据,从文件中读取

		time := gtime.Now().Format("His") // ServerChan 不能连续发送相同的消息,所以加上时间字符串 时时分分秒秒
		g.Log().Line().Info(chanId)
		push(chanId, "text"+time, "desp")
	})
}

func TestResumeRemind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		boot.InitDatabase()
		ResumeRemind()
	})
}
