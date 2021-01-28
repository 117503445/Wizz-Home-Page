package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

func NeedRole(role string) func(*ghttp.Request) {
	return func(r *ghttp.Request) {
		user := new(model.User)
		if err := r.GetCtxVar("user").Struct(user); err != nil {
			response.JsonOld(r, 401, "user not found")
		} else {
			if dao.HasRole(user, role) {
				r.Middleware.Next()
			} else {
				response.JsonOld(r, 401, "don't have role")
			}
		}
	}
}
