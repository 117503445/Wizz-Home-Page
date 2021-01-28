package router

import (
	"wizz-home-page/app/api"
	"wizz-home-page/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func middlewareAuth(r *ghttp.Request) {
	middleware.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)

		group.Group("/api", func(group *ghttp.RouterGroup) {

			group.Group("/stories", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Story.ReadAll)
				group.GET("/{id}", api.Story.ReadOne)
				group.POST("/", api.Story.Create)
				group.DELETE("/{id}", api.Story.Delete)
				group.PUT("/{id}", api.Story.Update)
			})

			group.POST("/auth/login", middleware.Auth.LoginHandler)

			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/", api.User.SignUp)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middlewareAuth)
					group.GET("/", api.User.GetInfo)
				})
			})
		})
	})
}
