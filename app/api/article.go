package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Article = new(articlesApi)

type articlesApi struct{}

// @Summary 获取所有文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ArticlesApiRep
// @Router /api/articles [get]
func (*articlesApi) ReadAll(r *ghttp.Request) {
	var articles []model.Articles
	if err := dao.Articles.Structs(&articles); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(articles) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var articlesRsp []model.ArticlesApiRep
		if err := gconv.Structs(articles, &articlesRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, articlesRsp)
	}
}

// @Summary 获取一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Success 200 {object} model.ArticlesApiRep
// @Failure 404 {string} string "{"message":"Article not found"}"
// @Router /api/articles/{id} [get]
func (*articlesApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var articles model.Articles
	if err := dao.Articles.Where("id = ", id).Struct(&articles); err != nil {
		response.JsonOld(r, 404, "")
	}
	var articleRsp model.ArticlesApiRep
	if err := gconv.Struct(articles, &articleRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.JsonOld(r, 200, articleRsp)
}

// @Summary 添加一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   articles      body model.ArticleApiCreateReq true  "文章"
// @Success 200 {object} model.ArticlesApiRep
// @Router /api/articles [POST]
// @Security JWT
func (*articlesApi) Create(r *ghttp.Request) {
	var (
		apiReq   *model.ArticleApiCreateReq
		articles *model.Articles
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if err := gconv.Struct(apiReq, &articles); err != nil {
		response.JsonOld(r, 400, "not a articles")
	}
	if result, err := dao.Articles.Insert(articles); err != nil {
		response.JsonOld(r, 500, "")
	} else {
		id, _ := result.LastInsertId()
		articles.Id = gconv.Int(id)

		var articleRsp model.ArticlesApiRep
		if err := gconv.Struct(articles, &articleRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, articleRsp)
	}
}

// @Summary 删除一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Article not found"}"
// @Router /api/articles/{id} [DELETE]
// @Security JWT
func (*articlesApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Articles.Where("id", id).Delete(); err != nil {
		response.JsonOld(r, 404, "")
	}
	response.JsonOld(r, 200, g.Map{"message": "delete success"})
}

// @Summary 更改一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Param   articles      body model.ArticleApiCreateReq true  "文章"
// @Success 200 {object} model.ArticlesApiRep
// @Failure 404 {string} string "{"message": "Article not found"}"
// @Router /api/articles/{id} [PUT]
// @Security JWT
func (*articlesApi) Update(r *ghttp.Request) {
	id := r.GetInt("id")
	var article model.Articles

	var (
		apiReq *model.ArticleApiCreateReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonOld(r, 400, "not a article")
	}
	if err := gconv.Struct(apiReq, &article); err != nil {
		response.JsonOld(r, 400, "not a article")
	}

	article.Id = id
	if _, err := dao.Articles.Data(article).Where("id", id).Update(); err != nil {
		response.JsonOld(r, 404, err.Error())
	} else {
		var articleRsp model.ArticlesApiRep
		if err := gconv.Struct(article, &articleRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, articleRsp)
	}
}
