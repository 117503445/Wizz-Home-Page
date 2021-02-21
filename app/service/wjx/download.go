package wjx

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/araddon/dateparse"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"strconv"
	"strings"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service"
)

// PullResume 拉取简历
func PullResume() {
	//isNew := DownloadExcel() // 下载 Excel
	//if isNew {          // 如果有新简历
	//	ParseExcel() // 解析简历,放入数据库
	//}
	DownloadExcel()
	ParseExcel()
}

// DownloadExcel 从问卷星下载 Excel 后与磁盘文件比较，如果有更新则写入磁盘并返回 True
func DownloadExcel() bool {
	cookie := g.MapStrStr{"SojumpSurvey": g.Cfg().GetString("wjx.SojumpSurvey")}

	header := g.MapStrStr{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"}

	activity := g.Cfg().GetString("wjx.activity")

	url := fmt.Sprintf("https://www.wjx.cn/wjx/activitystat/viewstatsummary.aspx?activity=%v&reportid=-1&dw=1&dt=0", activity)

	if r, err := g.Client().Cookie(cookie).Header(header).Get(url); err != nil {
		g.Log().Line().Panic(err)
		return false
	} else {
		defer r.Close()

		downloadExcel := r.ReadAll()

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

		// todo 可以不检测旧内容,直接写入,反正到时候根据ID判断是否重复

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
		//Todo 检测是否因为 SojumpSurvey 过期，导致下载的不是 Excel
		return
	}
	rows, err := f.GetRows("Sheet1")
	titleRow := rows[0]
	mapPropertyIndex := g.MapStrInt{}
	mapIndexProperty := g.MapIntStr{}
	for index, Property := range titleRow {
		mapPropertyIndex[Property] = index
		mapIndexProperty[index] = Property
	}

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[mapPropertyIndex["序号"]])

		g.DB().GetLogger().SetStdoutPrint(false)                           // 禁用 Count SQL 的 Std 输出
		if count, err := dao.Resumes.Where("id", id).Count(); err != nil { // todo 一次 查出所有 ID 放进列表里供判断
			g.Log().Line().Debug(err)
		} else {
			//g.Log().Line().Debug(count)
			if count > 0 {
				continue // id 存在,不插入
			}
		}
		g.DB().GetLogger().SetStdoutPrint(g.Cfg().GetBool("database.logger.Stdout"))

		name := row[mapPropertyIndex["姓名："]]

		sendTime, err := dateparse.ParseAny(row[mapPropertyIndex["提交答卷时间"]])
		if err != nil {
			g.Log().Line().Debug(err)
		}
		sendTimeStamp := sendTime.Unix()

		//experience :=  row[mapPropertyIndex["请在这里上传你的简历/过往项目或作品（大大加分项）"]]

		fileUrl := row[mapPropertyIndex["请在这里上传你的简历/过往项目或作品（大大加分项）"]]
		collegeMajor := row[mapPropertyIndex["学院及专业"]]

		MapDepartment := g.MapStrInt{"技术部-前端开发": 1, "技术部-后端开发": 2, "产品部": 3, "运营部-运营组": 4, "运营部-UI组": 4}
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

		experience := -1
		setExperienceProperty := gset.New(true)
		setExperienceProperty.Add("是否有开发项目的经历？", "是否有过互联网产品经历/项目经历/项目负责人经历/实习经历？（任一都可）", "是否有互联网运营经历/团队管理经历/新媒体运营或排版经历？（任一都可）", "是否有UI设计/海报设计经验？") // P V Z AD
		setExperienceProperty.Iterator(func(property interface{}) bool {
			value := row[mapPropertyIndex[property.(string)]]
			if strings.Contains(value, "跳过") {
				return true
			} else {
				if strings.Contains(value, "没有") {
					experience = 0
				} else {
					experience = 1
				}
				return false
			}
		})

		describe := "describe NOT FOUND"
		setDescribeProperty := gset.New(true)
		setDescribeProperty.Add("请描述一下你的过往经历", "请描述一下你的过往经历", "请描述一下你的过往经历", "请描述一下你的过往经历", "请介绍一下你的技能点/能力吧：）", "请描述一下你的项目开发经历", "请描述一下你的项目开发经历", "请描述一下你的项目/负责人/实习经历", "请描述一下你的运营/管理/新媒体运营经历") // T U Y AC AF R S W AA
		// g.Log().Line().Debug(setDescribeProperty)
		for index, property := range row {
			if setDescribeProperty.Contains(mapIndexProperty[index]) {
				if !strings.Contains(property, "跳过") {
					describe = property
				}
			}
		}

		interviewId := service.GetCurrentInterview().Id

		resume := &model.Resumes{
			Id:                     id,
			Describe:               describe,
			FileUrl:                fileUrl,
			CollegeMajor:           collegeMajor,
			Name:                   name,
			Gender:                 gender,
			Grade:                  grade,
			DepartmentType:         departmentType,
			Experience:             experience,
			TelephoneNumber:        MapContact["电话"],
			QqNumber:               MapContact["QQ"],
			WechatNumber:           MapContact["微信"],
			InterviewId:            interviewId,
			InterviewerId:          0,
			SendTime:               sendTimeStamp,
			InitialScreeningResult: 0,
			InitialScreeningTime:   0,
			InterviewResult:        0,
			InterviewEvaluation:    "",
			InterviewTime:          0,
		}

		service.DistributeInterviewers(resume)

		if _, err := dao.Resumes.Insert(resume); err != nil {
			g.Log().Line().Debug(err)
		}
	}
}
