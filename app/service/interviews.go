package service

import (
	"github.com/gogf/gf/frame/g"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func Count(interviews model.Interviews) (int, int, int) {
	SendNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Count()
	PassNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Where("interview_result", 1).Count()
	FailNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Where("interview_result", 2).Count()
	return SendNumber, PassNumber, FailNumber
}

// GetCurrentInterview 获得当前的面试项目.目前认为 最晚创建的面试项目 就是当前的面试项目
func GetCurrentInterview() *model.Interviews {
	if interview, err := dao.Interviews.Order("create_time desc").FindOne(); err != nil {
		g.Log().Line().Debug(err)
		return nil
	} else {
		// g.Log().Line().Debug(interview)
		return interview
	}
}
