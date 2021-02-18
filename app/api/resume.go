package api

import (
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service"
	"wizz-home-page/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Resume = new(resumesApi)

type resumesApi struct{}

// @Summary 获取搜索的分页简历
// @Tags 简历
// @Accept  json
// @Produce  json
// @Param   DepartmentType	query	int true  "面试部门"
// @Param   InterviewId	query	int true  "面试项目"
// @Param   InterviewResult	query	int true  "面试结果"
// @Param   Page	query	int true  "页数"
// @Success 200 {array} model.ResumesApiRep
// @Router /api/resumes [get]
func (*resumesApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	DepartmentType := r.GetInt("DepartmentType")
	InterviewId := r.GetInt("InterviewId")
	InterviewResult := r.GetInt("InterviewResult")
	Page := r.GetInt("Page")
	resumes, err := service.SortResumes(Page, InterviewId, DepartmentType, InterviewResult)
	if err != nil {
		g.Log().Line().Error(err)
		response.Json(r, response.Error, "", nil)
	}
	if resumes == nil {
		response.Json(r, response.ErrorNotExist, "", nil)
	}
	var resumesRsp []model.ResumesApiRep
	err = gconv.Structs(resumes.Content, &resumesRsp)
	if err != nil {
		g.Log().Line().Error(err)
	}
	resumes.Content = resumesRsp
	response.Json(r, response.Success, "", resumes)
}

// @Summary 获取一个简历
// @Tags 简历
// @Accept  json
// @Produce  json
// @Param   id      path int true  "简历id" default(1)
// @Success 200 {object} model.ResumesApiRep
// @Failure 404 {string} string "{"message":"Resume not found"}"
// @Router /api/resumes/{id} [get]
func (*resumesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var resumes model.Resumes
	if err := dao.Resumes.Where("id = ", id).Struct(&resumes); err != nil {
		response.JsonOld(r, 404, "")
	}
	var resumeRsp model.ResumesApiRep
	if err := gconv.Struct(resumes, &resumeRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, resumeRsp)
}

// @Summary 添加一个简历
// @Tags 简历
// @Accept  json
// @Produce  json
// @Param   resumes      body model.ResumeApiCreateReq true  "简历"
// @Success 200 {object} model.ResumesApiRep
// @Router /api/resumes [POST]
// @Security JWT
func (*resumesApi) Create(r *ghttp.Request) {
	var (
		apiReq  *model.ResumeApiCreateReq
		resumes *model.Resumes
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a resumes")
	}
	if err := gconv.Struct(apiReq, &resumes); err != nil {
		response.JsonOld(r, 400, "not a resumes")
	}
	if result, err := dao.Resumes.Insert(resumes); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		resumes.Id = gconv.Int(id)

		var resumeRsp model.ResumesApiRep
		if err := gconv.Struct(resumes, &resumeRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, resumeRsp)
	}
}

// @Summary 删除一个简历
// @Tags 简历
// @Accept  json
// @Produce  json
// @Param   id      path int true  "简历id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Resume not found"}"
// @Router /api/resumes/{id} [DELETE]
// @Security JWT
func (*resumesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Resumes.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个简历
// @Tags 简历
// @Accept  json
// @Produce  json
// @Param   id      path int true  "简历id" default(1)
// @Param   resumes      body model.ResumeApiCreateReq true  "简历"
// @Success 200 {object} model.ResumesApiRep
// @Failure 404 {string} string "{"message": "Resume not found"}"
// @Router /api/resumes/{id} [PUT]
// @Security JWT
func (*resumesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var resume model.Resumes

	var (
		apiReq *model.ResumeApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a resume")
	}
	if err := gconv.Struct(apiReq, &resume); err != nil {
		response.JsonOld(r, 400, "not a resume")
	}

	resume.Id = id
	if _, err := dao.Resumes.Data(resume).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var resumeRsp model.ResumesApiRep
		if err := gconv.Struct(resume, &resumeRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, resumeRsp)
	}
}