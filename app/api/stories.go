package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Story = new(storyApi)

type storyApi struct{}

// @Summary 获取所有历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Stories
// @Router /api/stories [get]
func (*storyApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var stories []model.Stories
	if err := dao.Stories.Structs(&stories); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(stories) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, stories)
	}
}

// @Summary 获取一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Success 200 {object} model.Stories
// @Failure 404 {string} string "{"message":"Story not found"}"
// @Router /api/stories/{id} [get]
func (*storyApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var story model.Stories
	if err := dao.Stories.Where("id = ", id).Struct(&story); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, story)
}

// @Summary 添加一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   story      body model.Stories true  "历史事件"
// @Success 200 {object} model.Stories
// @Router /api/stories [POST]
// @Security JWT
func (*storyApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.StoryApiCreateReq
		story  *model.Stories
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a story")
	}
	if err := gconv.Struct(apiReq, &story); err != nil {
		response.JsonOld(r, 400, "not a story")
	}
	if _, err := dao.Stories.Insert(story); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, story)
}

// @Summary 删除一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Story not found"}"
// @Router /api/stories/{id} [DELETE]
// @Security JWT
func (*storyApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Stories.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Param   story      body model.Stories true  "历史事件"
// @Success 200 {object} model.Stories
// @Failure 404 {string} string "{"message": "Story not found"}"
// @Router /api/stories/{id} [PUT]
// @Security JWT
func (*storyApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var story model.Stories

	var (
		apiReq *model.StoryApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a story")
	}
	if err := gconv.Struct(apiReq, &story); err != nil {
		response.JsonOld(r, 400, "not a story")
	}
	if _, err := dao.Stories.Data(story).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, story)
}
