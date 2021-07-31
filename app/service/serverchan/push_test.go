package serverchan

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestInterviewPush(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := "name.txt"
		if !gfile.Exists(path) {
			g.Log().Line().Error(path + " don't exist")
			t.Fail()
		}
		name := gfile.GetContents(path) // chanId 为敏感数据,从文件中读取

		title := fmt.Sprintf("%v简历", "后端")
		content := fmt.Sprintf("%v %v级 %v\n联系电话：%v\n微信：%v\nqq：%v\n\n%v\n\n<a href=\"%v\">简历下载链接</a>\n\n---\n\n请7天内联系投递者安排面试，或者告知他初筛未通过，7天内未填写则会触发重复提醒\n\n不需要面试也需要点击链接填写理由哦\n\n---\n\n<a href=\"%v\">点我前往面评填写页面</a>", "肖熙", "2019", "通信工程学院", "15777", "1234", "234215", "项目经历xxx", "https://baidu.com", "https://wizzstudio.com")

		Push(name, title, content)
	})
}
