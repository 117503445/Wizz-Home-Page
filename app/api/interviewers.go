package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service"
	"wizz-home-page/library/response"
)

var Interviewer = new(interviewersApi)

type interviewersApi struct{}

// @Summary 获取指定面试的所有面试官
// @Tags 面试官
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试项目id" default(1)
// @Success 200 {array} model.InterviewersApiRep
// @Router /api/interviewers/all/{id} [get]
func (*interviewersApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	id := r.GetInt("id")
	var interviewers []model.Interviewers
	if err := dao.Interviewers.Where("interview_id = ", id).Structs(&interviewers); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(interviewers) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var interviewersRsp []model.InterviewersApiRep
		if err := gconv.Structs(interviewers, &interviewersRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, interviewersRsp)
	}
}

// @Summary 获取一个面试官
// @Tags 面试官
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试官id" default(1)
// @Success 200 {object} model.InterviewersApiRep
// @Failure 404 {string} string "{"message":"Interviewer not found"}"
// @Router /api/interviewers/{id} [get]
func (*interviewersApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var interviewers model.Interviewers
	if err := dao.Interviewers.Where("id = ", id).Struct(&interviewers); err != nil {
		response.JsonOld(r, 404, "")
	}
	var interviewerRsp model.InterviewersApiRep
	if err := gconv.Struct(interviewers, &interviewerRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, interviewerRsp)
}

// @Summary 添加一个面试官
// @Tags 面试官
// @Accept  json
// @Produce  json
// @Param   interviewers      body model.InterviewerApiCreateReq true  "面试官"
// @Success 200 {object} model.InterviewersApiRep
// @Router /api/interviewers [POST]
// @Security JWT
func (*interviewersApi) Create(r *ghttp.Request) {
	var (
		apiReq       *model.InterviewerApiCreateReq
		interviewers *model.Interviewers
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a interviewers")
	}
	if err := gconv.Struct(apiReq, &interviewers); err != nil {
		response.JsonOld(r, 400, "not a interviewers")
	}
	if result, err := dao.Interviewers.Insert(interviewers); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		interviewers.Id = gconv.Int(id)

		service.SearchInterviewer(interviewers.InterviewId)

		var interviewerRsp model.InterviewersApiRep
		if err := gconv.Struct(interviewers, &interviewerRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, interviewerRsp)
	}
}

// @Summary 删除一个面试官
// @Tags 面试官
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试官id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Interviewer not found"}"
// @Router /api/interviewers/{id} [DELETE]
// @Security JWT
func (*interviewersApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Interviewers.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个面试官
// @Tags 面试官
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试官id" default(1)
// @Param   interviewers      body model.InterviewerApiCreateReq true  "面试官"
// @Success 200 {object} model.InterviewersApiRep
// @Failure 404 {string} string "{"message": "Interviewer not found"}"
// @Router /api/interviewers/{id} [PUT]
// @Security JWT
func (*interviewersApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var interviewer model.Interviewers

	var (
		apiReq *model.InterviewerApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a interviewer")
	}
	if err := gconv.Struct(apiReq, &interviewer); err != nil {
		response.JsonOld(r, 400, "not a interviewer")
	}

	interviewer.Id = id
	if _, err := dao.Interviewers.Data(interviewer).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var interviewerRsp model.InterviewersApiRep
		if err := gconv.Struct(interviewer, &interviewerRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, interviewerRsp)
	}
}
