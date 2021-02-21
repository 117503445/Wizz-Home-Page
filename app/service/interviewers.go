package service

import (
	"github.com/gogf/gf/frame/g"
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
	// todo 通知面试官
}
