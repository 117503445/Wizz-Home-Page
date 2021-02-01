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
		var membersRsp []model.MemberApiRep
		if err := gconv.Structs(members, &membersRsp); err != nil {
			g.Log().Line().Error(err)
		}

		//for m := range members{
		//	var mRsp model.MemberApiRep
		//	gconv.Structs()
		//}
		response.JsonOld(r, 200, membersRsp)
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
		apiReq  *model.MemberApiCreateReq
		members *model.Members
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

// @Summary 上移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} model.Members
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/up/{id} [PUT]
// @Security JWT
func (*membersApi) UpMember(r *ghttp.Request) {
	id := r.GetInt("id")
	membertype := r.GetQueryInt("type")
	var member model.Members

	if err := dao.Members.Where("id = ", id).Struct(&member); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	if membertype <= 0 && membertype > 4 {
		response.JsonOld(r, 404, "your type is not correct")
	} else if member.MemberType != membertype {
		response.JsonOld(r, 400, "your type do not match your id")
		return
	}
	var members []model.Members
	if err := dao.Members.Where("member_type = ", membertype).Structs(&members); err != nil {
		response.JsonOld(r, 404, err.Error())
	}
	if len(members) == 0 || len(members) == 1 {
		response.JsonOld(r, 400, "the number of member is 0 or 1")
	} else if members[0].Id == id {
		response.JsonOld(r, 400, "the first member can not up")
	}
	i := 1
	var m1 model.Members
	var m2 model.Members
	m1 = members[i]
	for m1.Id != id {
		i++
		m1 = members[i]
	}
	m2 = members[i-1]
	m1, m2 = m2, m1
	m1.Id, m2.Id = m2.Id, m1.Id
	dao.Members.Data(m1).Where("id", m1.Id).Update()
	dao.Members.Data(m2).Where("id", m2.Id).Update()
	dao.Members.Where("member_type", membertype).Structs(&members)
	response.JsonOld(r, 200, members)
}

// @Summary 下移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} model.Members
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /api/members/down/{id} [PUT]
// @Security JWT
func (*membersApi) DownMember(r *ghttp.Request) {
	id := r.GetInt("id")
	membertype := r.GetQueryInt("type")
	var member model.Members

	if err := dao.Members.Where("id = ", id).Struct(&member); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	if membertype <= 0 && membertype > 4 {
		response.JsonOld(r, 404, "your type is not correct")
	} else if member.MemberType != membertype {
		response.JsonOld(r, 400, "your type do not match your id")
		return
	}
	var members []model.Members
	if err := dao.Members.Where("member_type = ", membertype).Structs(&members); err != nil {
		response.JsonOld(r, 404, err.Error())
	}
	if len(members) == 0 || len(members) == 1 {
		response.JsonOld(r, 400, "the number of member is 0 or 1")
	} else if members[len(members)-1].Id == id {
		response.JsonOld(r, 400, "the last member can not up")
	}
	i := 0
	var m1 model.Members
	var m2 model.Members
	m1 = members[i]
	for m1.Id != id {
		i++
		m1 = members[i]
	}
	m2 = members[i+1]
	m1, m2 = m2, m1
	m1.Id, m2.Id = m2.Id, m1.Id
	dao.Members.Data(m1).Where("id", m1.Id).Update()
	dao.Members.Data(m2).Where("id", m2.Id).Update()
	dao.Members.Where("member_type", membertype).Structs(&members)
	response.JsonOld(r, 200, members)
}
