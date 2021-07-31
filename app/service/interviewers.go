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
// 返回是否成功分配
func DistributeInterviewers(resume *model.Resumes) bool {
	// 同部门面试官随机挑选
	interviewers, err := dao.Interviewers.Where("department_type", resume.DepartmentType).Where("interview_id", resume.InterviewId).FindAll()
	if err != nil {
		g.Log().Line().Error(err)
		return false
	}
	if interviewers == nil || len(interviewers) == 0 {
		//g.Log().Line().Error("interviewer NOT FOUND")
		//title := fmt.Sprintf("%v 部门 还未设置管理员", resume.DepartmentType)
		//content := fmt.Sprintf("%v 部门 还未设置管理员\n %v 的简历处理失败", resume.DepartmentType, resume.Name)
		//serverchan.Alarm(title, content)
		g.Log().Line().Error(fmt.Sprintf("%v 部门 还未设置管理员", resume.DepartmentType))
		return false
	}

	//此部门面试官中,寻找正在处理简历数最少的面试官
	minCount := 9999
	interviewer := &model.Interviewers{}

	for _, _interviewer := range interviewers {
		count, _ := dao.Resumes.Where("interviewer_id", _interviewer.Id).Count()
		if count < minCount {
			minCount = count
			interviewer = _interviewer
		}
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

	url := fmt.Sprintf("https://wizzstudio.com/#/pass?id=%v&jwt=%v", resume.Id, token) // todo load from config
	g.Log().Line().Debug(url)

	/*experienceStr := ""
	if resume.Experience == 0 {
		experienceStr = "无"
	} else {
		experienceStr = "有"
	}*/

	strFileUrl := fmt.Sprintf("<a href=\"%v\">简历下载链接</a>\n\n", resume.FileUrl)
	if resume.FileUrl == "" {
		strFileUrl = "" // 如果不存在简历,那么隐藏下载链接
	}

	departmentStr := ""

	switch {
	case resume.DepartmentType == 1:
		departmentStr = "前端"
	case resume.DepartmentType == 2:
		departmentStr = "产品"
	case resume.DepartmentType == 3:
		departmentStr = "后端"
	case resume.DepartmentType == 4:
		departmentStr = "运营/UI"
	default:
		g.Log().Line().Error(resume.DepartmentType)
	}
	title := fmt.Sprintf("%v简历", departmentStr)
	content := fmt.Sprintf("%v %v %v\n联系电话：%v\n微信：%v\nqq：%v\n\n%v\n\n%v---\n\n请7天内联系投递者安排面试，或者告知他初筛未通过，7天内未填写则会触发重复提醒\n\n不需要面试也需要点击链接填写理由哦\n\n---\n\n<a href=\"%v\">点我前往面评填写页面</a>", resume.Name, resume.Grade, resume.CollegeMajor, resume.TelephoneNumber, resume.WechatNumber, resume.QqNumber, resume.Describe, strFileUrl, url)
	serverchan.Push(interviewer.ServerchanId, title, content)
	return true
}

// 添加面试官后进行检索看面试项目的每个部门是否有面试官了
func SearchInterviewer(InterviewId int) {
	if search(InterviewId, 1) && search(InterviewId, 2) && search(InterviewId, 3) && search(InterviewId, 4) {
		var message *model.Messages
		err := dao.Messages.Where("resume_id", 0).Where("read_status", 0).Struct(&message)
		if err == nil {
			message.ReadStatus = 1
			if _, err = dao.Messages.Data(message).Where("id", message.Id).Update(); err != nil {
				g.Log().Line().Error(err)
			}
		}

	}
}

func search(InterviewId int, Type int) bool {
	n, err := dao.Interviewers.Where("interview_id", InterviewId).Where("department_type", Type).Count()
	if err != nil {
		g.Log().Line().Error(err)
	}
	return n != 0
}
