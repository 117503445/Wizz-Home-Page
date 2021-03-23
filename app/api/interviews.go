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

var Interview = new(interviewsApi)

type interviewsApi struct{}

// @Summary 获取所有面试项目
// @Tags 面试项目
// @Accept  json
// @Produce  json
// @Success 200 {array} model.InterviewsApiRep
// @Router /api/interviews [get]
func (*interviewsApi) ReadAll(r *ghttp.Request) {
	var interviews []model.Interviews
	if err := dao.Interviews.Structs(&interviews); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(interviews) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var interviewsRsp []model.InterviewsApiRep
		if err := gconv.Structs(interviews, &interviewsRsp); err != nil {
			g.Log().Line().Error(err)
		}
		for i, interview := range interviews {
			interviewsRsp[i].SendNumber, interviewsRsp[i].PassNumber, interviewsRsp[i].FailNumber = service.Count(interview)
		}
		response.JsonOld(r, 200, interviewsRsp)
	}
}

// @Summary 获取一个面试项目
// @Tags 面试项目
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试项目id" default(1)
// @Success 200 {object} model.InterviewsApiRep
// @Failure 404 {string} string "{"message":"Interview not found"}"
// @Router /api/interviews/{id} [get]
func (*interviewsApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var interviews model.Interviews
	if err := dao.Interviews.Where("id = ", id).Struct(&interviews); err != nil {
		response.JsonOld(r, 404, "")
	}
	var interviewRsp model.InterviewsApiRep
	if err := gconv.Struct(interviews, &interviewRsp); err != nil {
		g.Log().Line().Error(err)
	}
	interviewRsp.SendNumber, interviewRsp.PassNumber, interviewRsp.FailNumber = service.Count(interviews)
	response.JsonOld(r, 200, interviewRsp)
}

// @Summary 添加一个面试项目
// @Tags 面试项目
// @Accept  json
// @Produce  json
// @Param   interviews      body model.InterviewApiCreateReq true  "面试项目"
// @Success 200 {object} model.InterviewsApiRep
// @Router /api/interviews [POST]
// @Security JWT
func (*interviewsApi) Create(r *ghttp.Request) {
	var (
		apiReq     *model.InterviewApiCreateReq
		interviews *model.Interviews
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a interviews")
	}
	if err := gconv.Struct(apiReq, &interviews); err != nil {
		response.JsonOld(r, 400, "not a interviews")
	}
	if result, err := dao.Interviews.Insert(interviews); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		interviews.Id = gconv.Int(id)
		service.CreateMessage()
		var interviewRsp model.InterviewsApiRep
		if err := gconv.Struct(interviews, &interviewRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, interviewRsp)
	}
}

// @Summary 删除一个面试项目
// @Tags 面试项目
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试项目id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Interview not found"}"
// @Router /api/interviews/{id} [DELETE]
// @Security JWT
func (*interviewsApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Interviews.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个面试项目
// @Tags 面试项目
// @Accept  json
// @Produce  json
// @Param   id      path int true  "面试项目id" default(1)
// @Param   interviews      body model.InterviewApiCreateReq true  "面试项目"
// @Success 200 {object} model.InterviewsApiRep
// @Failure 404 {string} string "{"message": "Interview not found"}"
// @Router /api/interviews/{id} [PUT]
// @Security JWT
func (*interviewsApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var interview model.Interviews

	var (
		apiReq *model.InterviewApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a interview")
	}
	if err := gconv.Struct(apiReq, &interview); err != nil {
		response.JsonOld(r, 400, "not a interview")
	}

	interview.Id = id
	if _, err := dao.Interviews.Data(interview).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var interviewRsp model.InterviewsApiRep
		if err := gconv.Struct(interview, &interviewRsp); err != nil {
			g.Log().Line().Error(err)
		}
		interviewRsp.SendNumber, interviewRsp.PassNumber, interviewRsp.FailNumber = service.Count(interview)
		response.JsonOld(r, 200, interviewRsp)
	}
}
