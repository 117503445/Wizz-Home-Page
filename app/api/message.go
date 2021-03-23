package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Message = new(messagesApi)

type messagesApi struct{}

// @Summary 获取所有消息
// @Tags 消息
// @Accept  json
// @Produce  json
// @Success 200 {array} model.MessagesApiRep
// @Router /api/messages [get]
func (*messagesApi) ReadAll(r *ghttp.Request) {
	var messages []model.Messages
	if err := dao.Messages.Structs(&messages); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(messages) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var messagesRsp []model.MessagesApiRep
		if err := gconv.Structs(messages, &messagesRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, messagesRsp)
	}
}

// @Summary 获取一个消息
// @Tags 消息
// @Accept  json
// @Produce  json
// @Param   id      path int true  "消息id" default(1)
// @Success 200 {object} model.MessagesApiRep
// @Failure 404 {string} string "{"message":"Message not found"}"
// @Router /api/messages/{id} [get]
func (*messagesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var messages model.Messages
	if err := dao.Messages.Where("id = ", id).Struct(&messages); err != nil {
		response.JsonOld(r, 404, "")
	}
	var messageRsp model.MessagesApiRep
	if err := gconv.Struct(messages, &messageRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, messageRsp)
}

// @Summary 添加一个消息
// @Tags 消息
// @Accept  json
// @Produce  json
// @Param   messages      body model.MessageApiCreateReq true  "消息"
// @Success 200 {object} model.MessagesApiRep
// @Router /api/messages [POST]
// @Security JWT
func (*messagesApi) Create(r *ghttp.Request) {
	var (
		apiReq   *model.MessageApiCreateReq
		messages *model.Messages
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a messages")
	}
	if err := gconv.Struct(apiReq, &messages); err != nil {
		response.JsonOld(r, 400, "not a messages")
	}
	if result, err := dao.Messages.Insert(messages); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		messages.Id = gconv.Int(id)

		var messageRsp model.MessagesApiRep
		if err := gconv.Struct(messages, &messageRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, messageRsp)
	}
}

// @Summary 删除一个消息
// @Tags 消息
// @Accept  json
// @Produce  json
// @Param   id      path int true  "消息id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Message not found"}"
// @Router /api/messages/{id} [DELETE]
// @Security JWT
func (*messagesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Messages.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个消息
// @Tags 消息
// @Accept  json
// @Produce  json
// @Param   id      path int true  "消息id" default(1)
// @Param   messages      body model.MessageApiCreateReq true  "消息"
// @Success 200 {object} model.MessagesApiRep
// @Failure 404 {string} string "{"message": "Message not found"}"
// @Router /api/messages/{id} [PUT]
// @Security JWT
func (*messagesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var message model.Messages

	var (
		apiReq *model.MessageApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a message")
	}
	if err := gconv.Struct(apiReq, &message); err != nil {
		response.JsonOld(r, 400, "not a message")
	}

	message.Id = id
	if _, err := dao.Messages.Data(message).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var messageRsp model.MessagesApiRep
		if err := gconv.Struct(message, &messageRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, messageRsp)
	}
}
