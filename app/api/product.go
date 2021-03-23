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
// @Success 200 {array} model.ProductsApiRep
// @Router /api/products [get]
func (*productsApi) ReadAll(r *ghttp.Request) {
	var products []model.Products
	if err := dao.Products.Structs(&products); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(products) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var productsRsp []model.ProductsApiRep
		if err := gconv.Structs(products, &productsRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, productsRsp)
	}
}

// @Summary 获取一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Success 200 {object} model.ProductsApiRep
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
	var productRsp model.ProductsApiRep
	if err := gconv.Struct(products, &productRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, productRsp)
}

// @Summary 添加一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   products      body model.ProductApiCreateReq true  "产品"
// @Success 200 {object} model.ProductsApiRep
// @Router /api/products [POST]
// @Security JWT
func (*productsApi) Create(r *ghttp.Request) {
	var (
		apiReq   *model.ProductApiCreateReq
		products *model.Products
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if err := gconv.Struct(apiReq, &products); err != nil {
		response.JsonOld(r, 400, "not a products")
	}
	if result, err := dao.Products.Insert(products); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		products.Id = gconv.Int(id)

		var productRsp model.ProductsApiRep
		if err := gconv.Struct(products, &productRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, productRsp)
	}
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
// @Param   products      body model.ProductApiCreateReq true  "产品"
// @Success 200 {object} model.ProductsApiRep
// @Failure 404 {string} string "{"message": "Product not found"}"
// @Router /api/products/{id} [PUT]
// @Security JWT
func (*productsApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var product model.Products

	var (
		apiReq *model.ProductApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a product")
	}
	if err := gconv.Struct(apiReq, &product); err != nil {
		response.JsonOld(r, 400, "not a product")
	}

	product.Id = id
	if _, err := dao.Products.Data(product).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var productRsp model.ProductsApiRep
		if err := gconv.Struct(product, &productRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, productRsp)
	}
}
