package serverchan

import "github.com/gogf/gf/frame/g"

func Alarm(title string, content string) {
	chanIds := g.Cfg().GetArray("alarm.serverChans")
	for _, chanId := range chanIds {
		push(chanId.(string), title, content)
	}
}
