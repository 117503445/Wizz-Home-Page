package service

import (
	"errors"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func SortResumes(Page int, InterviewId int, DepartmentType int, InterviewResult int) ([]model.Resumes, error) {
	var resumes []model.Resumes
	err := dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Structs(&resumes)
	if err != nil {
		return nil, err
	}
	l := len(resumes)
	n1 := (Page - 1) * 10
	n2 := Page * 10
	if l <= n1 {
		err := errors.New("page incorrect")
		return nil, err
	}
	if l > n1 && l < n2 {
		return resumes[n1:], nil
	}
	return resumes[n1:n2], nil
}
