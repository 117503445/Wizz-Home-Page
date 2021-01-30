package wjx

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

// Download 从问卷星下载 Excel 后与磁盘文件比较，如果有更新则写入磁盘并返回 True
func Download() bool {
	cookie := g.MapStrStr{"SojumpSurvey": g.Cfg().GetString("wjx.SojumpSurvey")}

	header := g.MapStrStr{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"}

	activity := g.Cfg().GetString("wjx.activity")

	url := fmt.Sprintf("https://www.wjx.cn/wjx/activitystat/viewstatsummary.aspx?activity=%v&reportid=-1&dw=1&dt=0", activity)

	if r, err := g.Client().Cookie(cookie).Header(header).Get(url); err != nil {
		g.Log().Line().Panic(err)
		return false
	} else {
		defer r.Close()

		downloadExcel := r.ReadAll() // Todo 检测是否因为 SojumpSurvey 过期，导致下载的不是 Excel

		isNew := false

		excelPath := "./tmp/wjx.xlsx"
		if !gfile.Exists(excelPath) {
			isNew = true
		} else {
			diskExcel := gfile.GetBytes(excelPath)
			if !bytes.Equal(diskExcel, downloadExcel) {
				isNew = true
			}
		}

		g.Log().Line().Info(fmt.Sprintf("isNew: %v", isNew))

		if err := gfile.PutBytes(excelPath, downloadExcel); err != nil {
			g.Log().Line().Error(err)
		}

		return isNew
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
