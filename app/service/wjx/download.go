package wjx

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

// PullResume 拉起简历
func PullResume() {
	isNew := Download() // 下载 Excel
	if isNew {          // 如果有新简历
		ParseExcel() // 解析简历,放入数据库
	}
}

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

// ParseExcel 解析简历,根据姓名确定简历,如果简历不在数据库中,就创建简历
func ParseExcel() {
	f, err := excelize.OpenFile("./tmp/wjx.xlsx")
	if err != nil {
		return
	}
	rows, err := f.GetRows("Sheet1")
	titleRow := rows[0]
	mapPropertyIndex := g.MapStrInt{}
	for index, Property := range titleRow {
		g.Log().Line().Debug(Property, index)
		mapPropertyIndex[Property] = index
	}
	for _, row := range rows[1:] {
		name := row[mapPropertyIndex["姓名："]]
		id, _ := strconv.Atoi(row[mapPropertyIndex["序号"]])
		sendTime, _ := time.Parse("2006/01/02 03:04:05", row[mapPropertyIndex["提交答卷时间"]])
		sendTimeStamp := sendTime.Unix()
		//experience := row[mapPropertyIndex["是否有开发项目的经历？"]]
		//todo

		resume := &model.Resumes{
			Id:                     id,
			Describe:               "",
			FileUrl:                "",
			CollegeMajor:           "",
			Name:                   name,
			Gender:                 0,
			Grade:                  0,
			DepartmentType:         0,
			Experience:             0,
			InterviewId:            -1,
			InterviewerId:          -1,
			SendTime:               sendTimeStamp,
			InitialScreeningResult: 0,
			InitialScreeningTime:   0,
			InterviewResult:        0,
			InterviewEvaluation:    "",
			InterviewTime:          0,
		}
		if _, err := dao.Resumes.Insert(resume); err != nil {
			g.Log().Line().Debug(err)
		}
	}
}
