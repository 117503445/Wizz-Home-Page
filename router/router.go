package router

import (
	"wizz-home-page/app/api"
	"wizz-home-page/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS, middleware.ServerLog, middleware.HttpLog)

		group.ALL("/", api.Hello)

		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Group("/logs", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Log.ReadAll)
			})
			group.Group("/articles", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Article.ReadAll)
				group.GET("/{id}", api.Article.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Article.Create)
					group.DELETE("/{id}", api.Article.Delete)
					group.PUT("/{id}", api.Article.Update)
				})
			})
			group.Group("/image", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Image.ReadAll)
				group.GET("/read/{id}", api.Image.ReadOne)
				group.GET("/PlaceAndDomain", api.Image.GetPlaceAndDomain)
				group.GET("/BackGroundImageUrls", api.Image.GetBackGroundImageUrls)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Image.Create)
					group.DELETE("/{id}", api.Image.Delete)
					group.PUT("/{id}", api.Image.Update)
					group.GET("/UpToken", api.Image.GetUpToken)
				})
			})
			group.Group("/members", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Member.ReadAll)
				group.GET("/{id}", api.Member.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Member.Create)
					group.DELETE("/{id}", api.Member.Delete)
					group.PUT("/update/{id}", api.Member.Update)
					group.PUT("/up/{id}", api.Member.UpMember)
					group.PUT("/down/{id}", api.Member.DownMember)
				})
			})
			group.Group("/passage", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Passage.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Passage.Create)
					group.PUT("/", api.Passage.Update)
				})
			})
			group.Group("/products", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Product.ReadAll)
				group.GET("/{id}", api.Product.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Product.Create)
					group.DELETE("/{id}", api.Product.Delete)
					group.PUT("/{id}", api.Product.Update)
				})
			})
			group.Group("/stories", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Story.ReadAll)
				group.GET("/{id}", api.Story.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Story.Create)
					group.DELETE("/{id}", api.Story.Delete)
					group.PUT("/{id}", api.Story.Update)
				})
			})

			group.POST("/auth/login", middleware.Auth.LoginHandler)

			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/", api.User.SignUp)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin)
					group.GET("/", api.User.GetInfo)
				})
			})

			group.Group("/interviews", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Interview.ReadAll)
				group.GET("/{id}", api.Interview.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Interview.Create)
					group.DELETE("/{id}", api.Interview.Delete)
					group.PUT("/{id}", api.Interview.Update)
				})
			})

			group.Group("/interviewers", func(group *ghttp.RouterGroup) {
				group.GET("/all/{id}", api.Interviewer.ReadAll)
				group.GET("/{id}", api.Interviewer.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Interviewer.Create)
					group.DELETE("/{id}", api.Interviewer.Delete)
					group.PUT("/{id}", api.Interviewer.Update)
				})
			})

			group.Group("/resumes", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Resume.ReadAll)
				group.GET("/{id}", api.Resume.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Resume.Create)
					group.DELETE("/{id}", api.Resume.Delete)
					group.PUT("/{id}", api.Resume.Update)
				})
				group.Middleware(middleware.JWTLogin, middleware.NeedRole("interviewer"))
				group.PUT("/result/{id}", api.Resume.ResultUpdate)

			})

			group.Group("/messages", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Message.ReadAll)
				group.GET("/{id}", api.Message.ReadOne)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTLogin, middleware.NeedRole("admin"))
					group.POST("/", api.Message.Create)
					group.DELETE("/{id}", api.Message.Delete)
					group.PUT("/{id}", api.Message.Update)
				})
			})
		})
	})
}
