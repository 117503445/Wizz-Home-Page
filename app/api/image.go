package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Image = new(imagesApi)

type imagesApi struct{}

// @Summary 获取所有图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Images
// @Router /api/images [get]
func (*imagesApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var images []model.Images
	if err := dao.Images.Structs(&images); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(images) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, images)
	}
}

// @Summary 获取一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Success 200 {object} model.Images
// @Failure 404 {string} string "{"message":"Image not found"}"
// @Router /api/images/{id} [get]
func (*imagesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var images model.Images
	if err := dao.Images.Where("id = ", id).Struct(&images); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, images)
}

// @Summary 添加一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   images      body model.Images true  "图片"
// @Success 200 {object} model.Images
// @Router /api/images [POST]
// @Security JWT
func (*imagesApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.ImageApiCreateReq
		images  *model.Images
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if err := gconv.Struct(apiReq, &images); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if _, err := dao.Images.Insert(images); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, images)
}

// @Summary 删除一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Image not found"}"
// @Router /api/images/{id} [DELETE]
// @Security JWT
func (*imagesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Images.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Param   images      body model.Images true  "图片"
// @Success 200 {object} model.Images
// @Failure 404 {string} string "{"message": "Image not found"}"
// @Router /api/images/{id} [PUT]
// @Security JWT
func (*imagesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var images model.Images

	var (
		apiReq *model.ImageApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if err := gconv.Struct(apiReq, &images); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if _, err := dao.Images.Data(images).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, images)
}
