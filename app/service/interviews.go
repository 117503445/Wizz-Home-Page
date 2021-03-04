package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
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

// CreateMessage 创建面试后添加消息提醒管理员添加所有部门的面试官
func CreateMessage() {
	var message model.Messages
	message.ResumeId = 0
	message.ReadStatus = 0
	message.Content = "请为创建的最新项目添加面试官，否则无法进行下一步简历获取"
	message.CreateTime = gtime.TimestampMilli()
	if _, err := dao.Messages.Insert(message); err != nil {
		g.Log().Line().Error(err)
	}
}
