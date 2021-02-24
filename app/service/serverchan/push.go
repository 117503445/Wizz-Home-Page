package serverchan

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	url2 "net/url"
)

func Push(chanId string, text string, desp string) {
	encodeDesp := url2.QueryEscape(desp) // URL 编码
	url := fmt.Sprintf("https://sc.ftqq.com/%v.send?text=%v&desp=%v", chanId, text, encodeDesp)

	if _, err := g.Client().Get(url); err != nil {
		g.Log().Line().Error(err)
	}
}
