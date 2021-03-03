package service

import (
	"fmt"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

func PageResumes(Page int, InterviewId int, DepartmentType int, InterviewResult int) (*response.Page, error) {

	pageSize := 6
	n1 := (Page - 1) * pageSize

	totalSize, err := dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Count()
	if err != nil {
		return nil, err
	}

	var resumes []model.Resumes
	err = dao.Resumes.Where("interview_id", InterviewId).Where("department_type", DepartmentType).Where("interview_result", InterviewResult).Limit(n1, pageSize).Structs(&resumes)
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

func GetInterviewerName(id int) (string, error) {
	value, err := dao.Interviewers.Value("name", "id", id)
	if err != nil {
		return "", err
	}
	name := value.String()
	return name, nil
}

func ResultChange(result int) (int, error) {
	if result == 0 {
		return 2, nil
	} else if result == 1 {
		return 1, nil
	} else {
		err := fmt.Errorf("%s", "The resume result wrong")
		return 0, err
	}

}
