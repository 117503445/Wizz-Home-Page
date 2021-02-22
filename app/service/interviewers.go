package service

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
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
		g.Log().Line().Error("interviewer NOT FOUND") // todo 通知管理员
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

	// todo 通知面试官
}
