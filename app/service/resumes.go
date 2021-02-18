package service

import (
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func SortResumes(Page int, InterviewId int, DepartmentType int, InterviewResult int) ([]model.Resumes, error) {
	var resumes []model.Resumes
	n1 := (Page - 1) * 10
	n2 := Page * 10
	err := dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Limit(n1, n2).Structs(&resumes)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}
