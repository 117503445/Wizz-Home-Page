package middleware

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
)

func HttpLog(r *ghttp.Request) {
	r.Middleware.Next()

	url := r.URL
	statusCode := r.Response.Status
	requestMethod := r.Request.Method
	requestBody := ""
	if requestBodyBytes, err := ioutil.ReadAll(r.Request.Body); err == nil && len(requestBodyBytes) > 0 {
		requestBody = string(requestBodyBytes) + "\n"
	}
	responseBody := string(r.Response.Buffer())

	g.Log().Info(fmt.Sprintf("[HTTP] %v %v\n%v%v %v", requestMethod, url, requestBody, statusCode, responseBody))
}
