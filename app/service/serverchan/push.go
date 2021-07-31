package serverchan

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	url2 "net/url"
)

func Push(name string, title string, content string) {

	encodeText := url2.QueryEscape(title + "\n----\n" + content) // URL 编码
	url := fmt.Sprintf("https://push.gh.117503445.top:20000/push/text/v1?name=%v&text=%v", name, encodeText)

	if _, err := g.Client().Get(url); err != nil {
		g.Log().Line().Error(err)
	}
}
