package wjx

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
	"net/url"
	"strings"
)

func DownloadFile(fileUrl string) (string, error) {
	//fileUrl https://www.wjx.cn/wjx/viewfile.aspx?path=https%3a%2f%2ffile.sojump.cn%2f105939446_8_q11_1614223420ZbFzBc.pdf%3fattname%3d8_11_%25e6%25b5%2581%25e7%25a8%258b%25ef%25bc%2588%25e6%259b%25b4%25e6%2596%25b0%25e4%25ba%258e1.30.pdf&activity=105939446
	fileUrl, err := url.QueryUnescape(fileUrl)
	if err != nil {
		g.Log().Line().Debug(err)
		return "", err
	}
	//fileUrl https://www.wjx.cn/wjx/viewfile.aspx?path=https://file.sojump.cn/105939446_8_q11_1614223420ZbFzBc.pdf?attname=8_11_%e6%b5%81%e7%a8%8b%ef%bc%88%e6%9b%b4%e6%96%b0%e4%ba%8e1.30.pdf&activity=105939446
	fileUrl = strings.Split(strings.Split(fileUrl, "?")[1], "=")[1]
	//fileUrl https://file.sojump.cn/105939446_8_q11_1614223420ZbFzBc.pdf
	// g.Log().Line().Debug(fileUrl)

	fileName := strings.Split(fileUrl, "/")[len(strings.Split(fileUrl, "/"))-1]
	// fileName 105939446_8_q11_1614223420ZbFzBc.pdf
	// g.Log().Line().Debug(fileName)

	SojumpSurvey, err := gcache.Get("SojumpSurvey")
	if err != nil {
		g.Log().Line().Debug(err)
		return "", err
	}
	cookies := g.MapStrStr{"SojumpSurvey": SojumpSurvey.(string)}

	header := g.MapStrStr{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}

	params := g.MapStrStr{
		"path": fileUrl, "activity": g.Cfg().GetString("wjx.activity"), "down": "1"}

	file := g.Client().Cookie(cookies).Header(header).GetBytes("https://www.wjx.cn/wjx/viewfile.aspx", params)
	if err = gfile.PutBytes("./tmp/file/"+fileName, file); err != nil {
		g.Log().Line().Error(err)
		return "", err
	}

	return "/" + fileName, nil
}
