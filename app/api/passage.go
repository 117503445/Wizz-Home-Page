package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Passage = new(passagesApi)

type passagesApi struct{}

// @Summary 获取介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Passages
// @Failure 404 {string} string "{"message": "Passage not exist"}"
// @Router /api/passage [get]
func (*passagesApi) ReadOne(r *ghttp.Request) {
	//id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var passages model.Passages
	if err := dao.Passages.Where("id = 1").Struct(&passages); err != nil {
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
// @Router /api/passage [POST]
// @Security JWT
func (*passagesApi) Create(r *ghttp.Request) {
	var (
		apiReq   *model.PassageApiCreateReq
		passages *model.Passages
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if err := gconv.Struct(apiReq, &passages); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	passages.Id = 1
	if result, err := dao.Passages.Replace(passages); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		passages.Id = gconv.Int(id)
		response.JsonOld(r, 200, passages)
	}
}

// @Summary 更改一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   id      path int true  "介绍id" default(1)
// @Param   passages      body model.Passages true  "介绍"
// @Success 200 {object} model.Passages
// @Failure 404 {string} string "{"message": "Passage not found"}"
// @Router /api/passage [PUT]
// @Security JWT
func (*passagesApi) Update(r *ghttp.Request) {
	var (
		apiReq   *model.PassageApiCreateReq
		passages *model.Passages
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	if err := gconv.Struct(apiReq, &passages); err != nil {
		response.JsonOld(r, 400, "not a passages")
	}
	passages.Id = 1
	if result, err := dao.Passages.Replace(passages); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		passages.Id = gconv.Int(id)
		response.JsonOld(r, 200, passages)
	}
}
