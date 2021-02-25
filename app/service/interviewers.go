package service

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service/serverchan"
)

// DistributeInterviewers 分配面试官
func DistributeInterviewers(resume *model.Resumes) {
	// 同部门面试官随机挑选
	interviewer, err := dao.Interviewers.Where("department_type", resume.DepartmentType).Where("interview_id", resume.InterviewId).
		Order("rand()").FindOne()
	if err != nil {
		g.Log().Line().Error(err)
		return
	}

	if interviewer == nil {
		g.Log().Line().Error("interviewer NOT FOUND")
		title := fmt.Sprintf("%v 部门 还未设置管理员", resume.DepartmentType)
		content := fmt.Sprintf("%v 部门 还未设置管理员\n %v 的简历处理失败", resume.DepartmentType, resume.Name)
		serverchan.Alarm(title, content)
		return
	}
	resume.InterviewerId = interviewer.Id

	username := "interviewer"
	password := gfile.GetContents("./tmp/password/admin.txt")
	//todo read port in config
	response := g.Client().ContentJson().PostContent("http://localhost:80/api/auth/login", g.Map{
		"username": username,
		"password": password,
	})
	token := ""
	if js, err := gjson.DecodeToJson(response); err != nil {
		g.Log().Line().Error(err)
	} else {
		token = js.GetString("token")
	}

	url := fmt.Sprintf("https://wizzstudio.com/#/pass?id=%v&jwt=%v", resume.Id, token)
	g.Log().Line().Debug(url)

	experienceStr := ""
	if resume.Experience == 0 {
		experienceStr = "无"
	} else {
		experienceStr = "有"
	}

	title := fmt.Sprintf("%v的简历", resume.Name)
	content := fmt.Sprintf("%v %v %v\n（联系电话：%v）（微信：%v）（qq：%v）\n\n%v项目经历\n\n%v\n\n简历下载链接%v\n\n---\n\n请及时联系投递者安排面试，或者告知他初筛未通过哦\n\n面试结束后须返回该页面，点击链接填写面评，建议将该页面添加至浮窗\n\n不需要面试也需要点击链接填写理由哦\n\n---\n\n[点我前往面评填写页面](%v)", resume.Name, resume.Grade, resume.CollegeMajor, resume.TelephoneNumber, resume.WechatNumber, resume.QqNumber, experienceStr, resume.Describe, resume.FileUrl, url)
	serverchan.Push(interviewer.ServerchanId, title, content)
}
