package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Product = new(productsApi)

type productsApi struct{}

// @Summary 获取所有产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Products
// @Router /api/products [get]
func (*productsApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var products []model.Products
	if err := dao.Products.Structs(&products); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(products) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		response.JsonOld(r, 200, products)
	}
}

// @Summary 获取一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Success 200 {object} model.Products
// @Failure 404 {string} string "{"message":"Product not found"}"
// @Router /api/products/{id} [get]
func (*productsApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var products model.Products
	if err := dao.Products.Where("id = ", id).Struct(&products); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, products)
}

// @Summary 添加一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   products      body model.Products true  "产品"
// @Success 200 {object} model.Products
// @Router /api/products [POST]
// @Security JWT
func (*productsApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.ProductApiCreateReq
		products  *model.Products
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if err := gconv.Struct(apiReq, &products); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if _, err := dao.Products.Insert(products); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, products)
}

// @Summary 删除一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Product not found"}"
// @Router /api/products/{id} [DELETE]
// @Security JWT
func (*productsApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Products.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Param   products      body model.Products true  "产品"
// @Success 200 {object} model.Products
// @Failure 404 {string} string "{"message": "Product not found"}"
// @Router /api/products/{id} [PUT]
// @Security JWT
func (*productsApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var products model.Products

	var (
		apiReq *model.ProductApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if err := gconv.Struct(apiReq, &products); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if _, err := dao.Products.Data(products).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	}

	response.JsonOld(r, 200, products)
}
