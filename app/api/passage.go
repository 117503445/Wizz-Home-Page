package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Passage = new(passagesApi)

type passagesApi struct{}

// @Summary 获取所有介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Passages
// @Router /api/passages [get]
func (*passagesApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var passages []model.Passages
	if err := dao.Passages.Structs(&passages); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(passages) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, passages)
	}
}

// @Summary 获取一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   id      path int true  "介绍id" default(1)
// @Success 200 {object} model.Passages
// @Failure 404 {string} string "{"message":"Passage not found"}"
// @Router /api/passages/{id} [get]
func (*passagesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var passages model.Passages
	if err := dao.Passages.Where("id = ", id).Struct(&passages); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, passages)
}

// @Summary 添加一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   passages      body model.Passages true  "介绍"
// @Success 200 {object} model.Passages
// @Router /api/passages [POST]
// @Security JWT
func (*passagesApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.PassageApiCreateReq
		passages  *model.Passages
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if err := gconv.Struct(apiReq, &passages); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if _, err := dao.Passages.Insert(passages); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, passages)
}

// @Summary 删除一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   id      path int true  "介绍id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Passage not found"}"
// @Router /api/passages/{id} [DELETE]
// @Security JWT
func (*passagesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Passages.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   id      path int true  "介绍id" default(1)
// @Param   passages      body model.Passages true  "介绍"
// @Success 200 {object} model.Passages
// @Failure 404 {string} string "{"message": "Passage not found"}"
// @Router /api/passages/{id} [PUT]
// @Security JWT
func (*passagesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var passages model.Passages

	var (
		apiReq *model.PassageApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if err := gconv.Struct(apiReq, &passages); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if _, err := dao.Passages.Data(passages).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, passages)
}

