package serverchan

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func push(chanId string, text string, desp string) {
	url := fmt.Sprintf("https://sc.ftqq.com/%v.send?text=%v&desp=%v", chanId, text, desp)

	if _, err := g.Client().Get(url); err != nil {
		g.Log().Line().Error(err)
	}
}
