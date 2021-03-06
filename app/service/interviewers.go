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
		count, _ := dao.Resumes.Where("interviewer_id", _interviewer.Id).Where("interview_result", 0).Count()
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

	experienceStr := ""
	if resume.Experience == 0 {
		experienceStr = "无"
	} else {
		experienceStr = "有"
	}

	title := fmt.Sprintf("%v的简历", resume.Name)
	content := fmt.Sprintf("%v %v %v\n（联系电话：%v）（微信：%v）（qq：%v）\n\n%v项目经历\n\n%v\n\n[简历下载链接](%v)\n\n---\n\n请及时联系投递者安排面试，或者告知他初筛未通过哦\n\n面试结束后须返回该页面，点击链接填写面评，建议将该页面添加至浮窗\n\n不需要面试也需要点击链接填写理由哦\n\n---\n\n[点我前往面评填写页面](%v)", resume.Name, resume.Grade, resume.CollegeMajor, resume.TelephoneNumber, resume.WechatNumber, resume.QqNumber, experienceStr, resume.Describe, resume.FileUrl, url)
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
