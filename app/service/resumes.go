package service

import (
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

func SortResumes(Page int, InterviewId int, DepartmentType int, InterviewResult int) (*response.Page, error) {

	pageSize := 10
	n1 := (Page - 1) * pageSize
	n2 := Page * pageSize

	totalSize, err := dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Count()
	if err != nil {
		return nil, err
	}

	var resumes []model.Resumes
	err = dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Limit(n1, n2).Structs(&resumes)
	if err != nil {
		return nil, err
	}

	p := response.Page{
		PageNum:    Page,
		TotalPages: totalSize/pageSize + 1,
		PageSize:   pageSize,
		TotalSize:  totalSize,
		Content:    resumes,
	}
	return &p, nil
}
