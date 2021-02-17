package service

import (
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func Count(interviews model.Interviews) (int, int, int) {
	SendNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Count()
	PassNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Where("interview_result", 1).Count()
	FailNumber, _ := dao.Resumes.Where("interview_id", interviews.Id).Where("interview_result", 2).Count()
	return SendNumber, PassNumber, FailNumber
}
