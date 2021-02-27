package wjx

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		GetSojumpSurvey()
		fileUrl, err := DownloadFile("https://www.wjx.cn/wjx/viewfile.aspx?path=https%3a%2f%2ffile.sojump.cn%2f105939446_8_q11_1614223420ZbFzBc.pdf%3fattname%3d8_11_%25e6%25b5%2581%25e7%25a8%258b%25ef%25bc%2588%25e6%259b%25b4%25e6%2596%25b0%25e4%25ba%258e1.30.pdf&activity=105939446")
		if err != nil {
			g.Log().Line().Debug(err)
		} else {
			g.Log().Line().Debug(fileUrl)
		}
	})
}
