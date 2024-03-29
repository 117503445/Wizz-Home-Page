package service

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
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

// GetEvaluationURL 获取面评填写 URL
func GetEvaluationURL(resumeID int) string {
	username := "interviewer"
	password := gfile.GetContents("./tmp/password/admin.txt")

	loginURL := fmt.Sprintf("http://localhost%v/api/auth/login", g.Cfg().GetString("server.Address"))
	g.Log().Line().Debug(loginURL)
	content := g.Client().ContentJson().PostContent(loginURL, g.Map{
		"username": username,
		"password": password,
	})
	token := ""
	if js, err := gjson.DecodeToJson(content); err != nil {
		g.Log().Line().Error(err)
	} else {
		token = js.GetString("token")
	}
	url := fmt.Sprintf("https://wizzstudio.com/#/pass?id=%v&jwt=%v", resumeID, token) // todo load from config
	return url
}

func RepeatResume(resume *model.Resumes) (Repeated bool, interviewerId int) {
	count, err := dao.Resumes.Where("name", resume.Name).Where("department_type", resume.DepartmentType).Where("college_major", resume.CollegeMajor).Where("interview_result = 0").Count()
	if err != nil {
		g.Log().Line().Debug(err)
		return false, 0
	}
	// 无重复简历
	if count == 0 {
		return false, 0
	}
	// 有重复简历
	valueId, err := dao.Resumes.Where("name", resume.Name).Where("department_type", resume.DepartmentType).Where("college_major", resume.CollegeMajor).Where("interview_result = 0").Value("id")
	if err != nil {
		g.Log().Line().Debug(err)
		return false, 0
	}
	id := valueId.Int()

	valueInterviewerId, err := dao.Resumes.Value("interviewer_id", "id", id)
	if err != nil {
		g.Log().Line().Debug(err)
		return false, 0
	}
	// 旧简历状态改为无效
	_, err = dao.Resumes.Data("interview_result = 3").Where("id", id).Update()
	if err != nil {
		g.Log().Line().Debug(err)
		return false, 0
	}
	// 返回老简历的面试官Id
	return true, valueInterviewerId.Int()

}
