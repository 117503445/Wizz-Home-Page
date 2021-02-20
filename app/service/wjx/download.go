package wjx

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/araddon/dateparse"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"strconv"
	"strings"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
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

		// g.Log().Line().Info(fmt.Sprintf("isNew: %v", isNew))

		if err := gfile.PutBytes(excelPath, downloadExcel); err != nil {
			g.Log().Line().Error(err)
		}

		return isNew
	}
}

// ParseExcel 解析简历,根据姓名确定简历,如果简历不在数据库中,就插入简历
func ParseExcel() {
	f, err := excelize.OpenFile("./tmp/wjx.xlsx")
	if err != nil {
		return
	}
	rows, err := f.GetRows("Sheet1")
	titleRow := rows[0]
	mapPropertyIndex := g.MapStrInt{}
	for index, Property := range titleRow {
		mapPropertyIndex[Property] = index
	}
	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[mapPropertyIndex["序号"]])

		if count, err := dao.Resumes.Where("id", id).Count(); err != nil {
			g.Log().Line().Debug(err)
		} else {
			//g.Log().Line().Debug(count)
			if count > 0 {
				continue // id 存在,不插入
			}
		}

		name := row[mapPropertyIndex["姓名："]]

		sendTime, err := dateparse.ParseAny(row[mapPropertyIndex["提交答卷时间"]])
		if err != nil {
			g.Log().Line().Debug(err)
		}
		sendTimeStamp := sendTime.Unix()

		//experience :=  row[mapPropertyIndex["请在这里上传你的简历/过往项目或作品（大大加分项）"]]

		fileUrl := row[mapPropertyIndex["请在这里上传你的简历/过往项目或作品（大大加分项）"]]
		collegeMajor := row[mapPropertyIndex["学院及专业"]]

		MapDepartment := g.MapStrInt{"技术部-前端开发": 1, "技术部-后端开发": 2, "产品部": 3, "运营部-运营组": 4, "运营部-UI组": 5}
		departmentType := MapDepartment[row[mapPropertyIndex["意向加入的部门"]]]

		MapGender := g.MapStrInt{"男": 1, "女": 0}
		gender := MapGender[row[mapPropertyIndex["性别："]]]

		gradeStr := row[mapPropertyIndex["年级"]] //"2020级"
		gradeStr = gradeStr[:len(gradeStr)-3]   //"2020"
		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			g.Log().Line().Debug(err)
		}

		MapContact := g.MapStrStr{}
		contactsStr := row[mapPropertyIndex["联系方式"]]
		for _, contactStr := range strings.Split(contactsStr, "┋") {
			splitIndex := strings.Index(contactStr, "〖")
			key := contactStr[:splitIndex]
			value := contactStr[splitIndex+3 : len(contactStr)-3]
			MapContact[key] = value
		}
		// g.Log().Line().Debug(MapContact)

		resume := &model.Resumes{
			Id:                     id,
			Describe:               "等会再写Describe", //todo
			FileUrl:                fileUrl,
			CollegeMajor:           collegeMajor,
			Name:                   name,
			Gender:                 gender,
			Grade:                  grade,
			DepartmentType:         departmentType,
			Experience:             0,
			TelephoneNumber:        MapContact["电话"],
			QqNumber:               MapContact["QQ"],
			WechatNumber:           MapContact["微信"],
			InterviewId:            0,
			InterviewerId:          0,
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
