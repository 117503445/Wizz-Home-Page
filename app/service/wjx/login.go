package wjx

import (
	"github.com/antchfx/htmlquery"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"net/http"
	"strings"
	"time"
)

// 模拟问卷星登录,获得 SojumpSurvey
// 多次调用需间隔 60 s
func GetSojumpSurvey() {
	if contains, _ := gcache.Contains("login"); contains {
		g.Log().Line().Debug("last login in 60s, reject this login :(")
		return
	}
	gcache.Set("Login", true, time.Minute)

	response, err := g.Client().Get("https://www.wjx.cn/css/index.css")
	if err != nil {
		g.Log().Line().Panic(err)
	}
	SERVERID := ""
	for _, cookie := range response.Cookies() {
		if cookie.Name == "SERVERID" {
			SERVERID = cookie.Value
		}
	}
	// g.Log().Line().Debug(SERVERID)

	__VIEWSTATE := ""
	__EVENTVALIDATION := ""

	header := g.MapStrStr{
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"Origin":                    "https://www.wjx.cn",
		"Content-Type":              "application/x-www-form-urlencoded",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Referer":                   "https://www.wjx.cn/login.aspx",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8"}

	text := g.Client().Header(header).GetContent("https://www.wjx.cn/login.aspx")
	doc, err := htmlquery.Parse(strings.NewReader(text))

	node := htmlquery.FindOne(doc, "//*[@id=\"__VIEWSTATE\"]")
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			__VIEWSTATE = attr.Val
		}
	}

	node = htmlquery.FindOne(doc, "//*[@id=\"__EVENTVALIDATION\"]")
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			__EVENTVALIDATION = attr.Val
		}
	}

	// g.Log().Line().Debug(__VIEWSTATE)
	// g.Log().Line().Debug(__EVENTVALIDATION)

	username := g.Cfg().GetString("wjx.username")
	password := g.Cfg().GetString("wjx.password")

	cookies := g.MapStrStr{"SERVERID": SERVERID}
	data := g.MapStrStr{"__VIEWSTATE": __VIEWSTATE,
		"__EVENTVALIDATION": __EVENTVALIDATION,
		"UserName":          username,
		"Password":          password,
		"LoginButton":       "登 录"}
	c := g.Client()
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse // 禁用跳转
	}
	response, err = c.Header(header).Cookie(cookies).Post("https://www.wjx.cn/login.aspx", data)
	if err != nil {
		g.Log().Line().Panic(err)
	}
	SojumpSurvey := response.GetCookie("SojumpSurvey")
	gcache.Set("SojumpSurvey", SojumpSurvey, 0)
	g.Log().Line().Debug(SojumpSurvey)
}
