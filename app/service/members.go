package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

func MembersChangePosition(r *ghttp.Request, isUp bool) {
	id := r.GetInt("id")
	memberType := r.GetQueryInt("type")
	var member model.Members

	if err := dao.Members.Where("id = ", id).Struct(&member); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	if memberType <= 0 && memberType > 4 {
		response.JsonOld(r, 404, "your type is not correct")
	} else if member.MemberType != memberType {
		response.JsonOld(r, 400, "your type do not match your id")
		return
	}
	var members []model.Members
	if err := dao.Members.Where("member_type = ", memberType).Structs(&members); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	if isUp {
		if len(members) == 0 || len(members) == 1 {
			response.JsonOld(r, 400, "the number of member is 0 or 1")
		} else if members[0].Id == id {
			response.JsonOld(r, 400, "the first member can not up")
		}
	} else {
		if len(members) == 0 || len(members) == 1 {
			response.JsonOld(r, 400, "the number of member is 0 or 1")
		} else if members[len(members)-1].Id == id {
			response.JsonOld(r, 400, "the last member can not down")
		}
	}

	var i int
	if isUp {
		i = 1
	} else {
		i = 0
	}

	var m1 model.Members
	var m2 model.Members
	m1 = members[i]
	for m1.Id != id {
		i++
		m1 = members[i]
	}

	if isUp {
		m2 = members[i-1]
	} else {
		m2 = members[i+1]
	}

	m1, m2 = m2, m1
	m1.Id, m2.Id = m2.Id, m1.Id
	dao.Members.Data(m1).Where("id", m1.Id).Update()
	dao.Members.Data(m2).Where("id", m2.Id).Update()
	dao.Members.Where("member_type", memberType).Structs(&members)

	var membersRsp []model.MembersApiRep
	if err := gconv.Structs(members, &membersRsp); err != nil {
		g.Log().Line().Error(err)
	}

	response.JsonOld(r, 200, membersRsp)
}
