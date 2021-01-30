package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Member = new(membersApi)

type membersApi struct{}

// @Summary 获取所有成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Members
// @Router /api/members [get]
func (*membersApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var members []model.Members
	if err := dao.Members.Structs(&members); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(members) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, members)
	}
}

// @Summary 获取一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Success 200 {object} model.Members
// @Failure 404 {string} string "{"message":"Member not found"}"
// @Router /api/members/{id} [get]
func (*membersApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var members model.Members
	if err := dao.Members.Where("id = ", id).Struct(&members); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, members)
}

// @Summary 添加一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   members      body model.Members true  "成员"
// @Success 200 {object} model.Members
// @Router /api/members [POST]
// @Security JWT
func (*membersApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.MemberApiCreateReq
		members  *model.Members
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if err := gconv.Struct(apiReq, &members); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if _, err := dao.Members.Insert(members); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, members)
}

// @Summary 删除一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/{id} [DELETE]
// @Security JWT
func (*membersApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Members.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   members      body model.Members true  "成员"
// @Success 200 {object} model.Members
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/{id} [PUT]
// @Security JWT
func (*membersApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var members model.Members

	var (
		apiReq *model.MemberApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if err := gconv.Struct(apiReq, &members); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if _, err := dao.Members.Data(members).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, members)
}
