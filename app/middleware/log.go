package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
	"strings"
	"time"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func ServerLog(r *ghttp.Request) {

	r.Middleware.Next()
	url := r.URL
	statusCode := r.Response.Status
	requestMethod := r.Request.Method

	userName := ""
	user := new(model.User)
	if err := r.GetCtxVar("user").Struct(user); err == nil {
		userName = user.Username
	}
	if userName == "" {
		userName = "Guest"
	}

	if requestMethod == "POST" || requestMethod == "PUT" || requestMethod == "DELETE" {
		if _, err := dao.Logs.Insert(model.Logs{
			TimeStamp:     time.Now().Unix(),
			RequestMethod: requestMethod,
			Username:      userName,
			RequestUrl:    url.Path,
			ResponseCode:  strconv.Itoa(statusCode),
			ModelName:     strings.Split(url.Path, "/")[2],
		}); err != nil {
			g.Log().Line().Error(err)
		}
	}

}
