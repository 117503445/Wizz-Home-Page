package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Story = new(storiesApi)

type storiesApi struct{}

// @Summary 获取所有历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Success 200 {array} model.StoriesApiRep
// @Router /api/stories [get]
func (*storiesApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var stories []model.Stories
	if err := dao.Stories.Structs(&stories); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(stories) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var storiesRsp []model.StoriesApiRep
		if err := gconv.Structs(stories, &storiesRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, storiesRsp)
	}
}

// @Summary 获取一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Success 200 {object} model.StoriesApiRep
// @Failure 404 {string} string "{"message":"Story not found"}"
// @Router /api/stories/{id} [get]
func (*storiesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var stories model.Stories
	if err := dao.Stories.Where("id = ", id).Struct(&stories); err != nil {
		response.JsonOld(r, 404, "")
	}
	var storyRsp model.StoriesApiRep
	if err := gconv.Struct(stories, &storyRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, storyRsp)
}

// @Summary 添加一个历史事件
// @Tags 历史事件
// @Accept  json
// @Produce  json
// @Param   stories      body model.StoryApiCreateReq true  "历史事件"
// @Success 200 {object} model.StoriesApiRep
// @Router /api/stories [POST]
// @Security JWT
func (*storiesApi) Create(r *ghttp.Request) {
	var (
		apiReq  *model.StoryApiCreateReq
		stories *model.Stories
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a stories")
	}
	if err := gconv.Struct(apiReq, &stories); err != nil {
		response.JsonOld(r, 400, "not a stories")
	}
	if result, err := dao.Stories.Insert(stories); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		stories.Id = gconv.Int(id)

		var storyRsp model.StoriesApiRep
		if err := gconv.Struct(stories, &storyRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, storyRsp)
	}
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
func (*storiesApi) Delete(r *ghttp.Request) {
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
// @Param   stories      body model.StoryApiCreateReq true  "历史事件"
// @Success 200 {object} model.StoriesApiRep
// @Failure 404 {string} string "{"message": "Story not found"}"
// @Router /api/stories/{id} [PUT]
// @Security JWT
func (*storiesApi) Update(r *ghttp.Request) {
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

	story.Id = id
	if _, err := dao.Stories.Data(story).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var storyRsp model.StoriesApiRep
		if err := gconv.Struct(story, &storyRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, storyRsp)
	}
}
