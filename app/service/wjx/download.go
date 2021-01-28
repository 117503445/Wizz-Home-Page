package wjx

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

func Download() {
	cookie := g.MapStrStr{"SojumpSurvey": g.Cfg().GetString("wjx.SojumpSurvey")}

	header := g.MapStrStr{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"}

	activity := g.Cfg().GetString("wjx.activity")

	url := fmt.Sprintf("https://www.wjx.cn/wjx/activitystat/viewstatsummary.aspx?activity=%v&reportid=-1&dw=1&dt=2", activity)

	if r, err := g.Client().Cookie(cookie).Header(header).Get(url); err != nil {
		g.Log().Line().Error(err)
	} else {
		defer r.Close()
		if err := gfile.PutBytes("./tmp/wjx.xlsx", r.ReadAll()); err != nil {
			g.Log().Line().Error(err)
		}
	}
}

func ParseExcel() {
	f, err := excelize.OpenFile("./tmp/wjx.xlsx")
	if err != nil {
		g.Log().Line().Error(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
