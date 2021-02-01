package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service/qiniu"
	"wizz-home-page/library/response"
)

var Image = new(imagesApi)

type imagesApi struct{}

// @Summary 获取所有图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ImagesApiRep
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
		var imagesRsp []model.ImagesApiRep
		if err := gconv.Structs(images, &imagesRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, imagesRsp)
	}
}

// @Summary 获取一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   id      path int true  "图片id" default(1)
// @Success 200 {object} model.ImagesApiRep
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
	var imageRsp model.ImagesApiRep
	if err := gconv.Struct(images, &imageRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, imageRsp)
}

// @Summary 添加一个图片
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   images      body model.ImageApiCreateReq true  "图片"
// @Success 200 {object} model.ImagesApiRep
// @Router /api/images [POST]
// @Security JWT
func (*imagesApi) Create(r *ghttp.Request) {
	var (
		apiReq  *model.ImageApiCreateReq
		images *model.Images
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if err := gconv.Struct(apiReq, &images); err != nil {
		response.JsonOld(r, 400, "not a images")
	}
	if result, err := dao.Images.Insert(images); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		images.Id = gconv.Int(id)

		var imageRsp model.ImagesApiRep
		if err := gconv.Struct(images, &imageRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, imageRsp)
	}
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
// @Param   images      body model.ImageApiCreateReq true  "图片"
// @Success 200 {object} model.ImagesApiRep
// @Failure 404 {string} string "{"message": "Image not found"}"
// @Router /api/images/{id} [PUT]
// @Security JWT
func (*imagesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var image model.Images

	var (
		apiReq *model.ImageApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a image")
	}
	if err := gconv.Struct(apiReq, &image); err != nil {
		response.JsonOld(r, 400, "not a image")
	}

	image.Id = id
	if _, err := dao.Images.Data(image).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var imageRsp model.ImagesApiRep
		if err := gconv.Struct(image, &imageRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, imageRsp)
	}
}

// @Summary 获取七牛云空间的地区和域名
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string "{"domain":"q52qkptnh.bkt.clouddn.com","place": "华东}"
// @Router /api/image/PlaceAndDomain [get]
func (*imagesApi) GetPlaceAndDomain(r *ghttp.Request) {
	place := g.Cfg().GetString("qiniu.place")
	domain := g.Cfg().GetString("qiniu.domain")

	response.JsonOld(r, 200, g.Map{"place": place, "domain": domain})
}

// @Summary 获取一个上传文件的upToken
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   fileName      query string true  "要上传的文件名,如 abc.png"
// @Success 200 {string} string "upToken"
// @Router /api/image/UpToken [get]
// @Security JWT
func (*imagesApi) GetUpToken(r *ghttp.Request) {
	fileName := r.GetQueryString("fileName")
	upToken := qiniu.GetUpToken(fileName)
	response.JsonOld(r, 200, "\""+upToken+"\"")
}

// @Summary 获取所有背景图片的url
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string "["http://q52qkptnh.bkt.clouddn.com/1.png","http://q52qkptnh.bkt.clouddn.com/2.png","http://q52qkptnh.bkt.clouddn.com/3.png"]"
// @Router /api/image/BackGroundImageUrls [get]
func (*imagesApi) GetBackGroundImageUrls(r *ghttp.Request) {
	urls := g.Cfg().GetArray("background.backgroundImageUrls")
	// g.Log().Line().Debug(urls)
	response.JsonOld(r, 200, urls)
}
