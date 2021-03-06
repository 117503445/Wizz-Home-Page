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

var Member = new(membersApi)

type membersApi struct{}

// @Summary 获取所有成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Success 200 {array} model.MembersApiRep
// @Router /api/members [get]
func (*membersApi) ReadAll(r *ghttp.Request) {
	var members []model.Members
	if err := dao.Members.Structs(&members); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(members) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var membersRsp []model.MembersApiRep
		if err := gconv.Structs(members, &membersRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, membersRsp)
	}
}

// @Summary 获取一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Success 200 {object} model.MembersApiRep
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
	var memberRsp model.MembersApiRep
	if err := gconv.Struct(members, &memberRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, memberRsp)
}

// @Summary 添加一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   members      body model.MemberApiCreateReq true  "成员"
// @Success 200 {object} model.MembersApiRep
// @Router /api/members [POST]
// @Security JWT
func (*membersApi) Create(r *ghttp.Request) {
	var (
		apiReq  *model.MemberApiCreateReq
		members *model.Members
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if err := gconv.Struct(apiReq, &members); err != nil {
		response.JsonOld(r, 400, "not a members")
	}
	if result, err := dao.Members.Insert(members); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		members.Id = gconv.Int(id)

		var memberRsp model.MembersApiRep
		if err := gconv.Struct(members, &memberRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, memberRsp)
	}
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
// @Param   members      body model.MemberApiCreateReq true  "成员"
// @Success 200 {object} model.MembersApiRep
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/update/{id} [PUT]
// @Security JWT
func (*membersApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var member model.Members

	var (
		apiReq *model.MemberApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a member")
	}
	if err := gconv.Struct(apiReq, &member); err != nil {
		response.JsonOld(r, 400, "not a member")
	}

	member.Id = id
	if _, err := dao.Members.Data(member).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var memberRsp model.MembersApiRep
		if err := gconv.Struct(member, &memberRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, memberRsp)
	}
}

// @Summary 上移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} model.MembersApiRep
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/up/{id} [PUT]
// @Security JWT
func (*membersApi) UpMember(r *ghttp.Request) {
	service.MembersChangePosition(r, true)
}

// @Summary 下移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} model.MembersApiRep
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/down/{id} [PUT]
// @Security JWT
func (*membersApi) DownMember(r *ghttp.Request) {
	service.MembersChangePosition(r, false)
}
