package Middlewares

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

// 封装 gin.ResponseWriter, gin.Context 进行 Write 时也对 body 进行了 Write 操作
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type nameModel struct {
	Name string
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	// 长度小于1024才进行写入
	if len(b) <= 1024 && r.body.Len() <= 1024 {
		r.body.Write(b)
	}
	return r.ResponseWriter.Write(b)
}
func GetLogMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		timestamp := time.Now().Unix()

		id := c.Params.ByName("id")
		method := c.Request.Method
		name := "Can't find name"
		uri := c.Request.RequestURI

		if method == "PUT" || method == "DELETE" {


			if strings.Contains(uri, "stories") {
				var story models.Story
				Global.Database.First(&story, id)
				name = story.Name
			} else if strings.Contains(uri, "products") {
				var product models.Product
				Global.Database.First(&product, id)
				name = product.Name
			} else if strings.Contains(uri, "members") {
				var member models.Member
				Global.Database.First(&member, id)
				name = member.Name
			}
		} else if method == "POST" {
			var nm nameModel
			_ = c.ShouldBind(&nm)
			name = nm.Name
		}

		//t := time.Now()
		// before request
		c.Next()
		// after request
		//latency := time.Since(t)
		//log.Println(c.ClientIP())
		//log.Println(latency)
		//log.Println(c.Request.Method)
		//log.Println(c.Request.Header)
		user, exists := c.Get(identityKey)
		if exists {

			serverLog := models.ServerLog{
				TimeStamp:     timestamp,
				RequestMethod: c.Request.Method,
				RequestURI:    uri,
				ResponseCode:  strconv.Itoa(c.Writer.Status()),
				Username:      user.(*User).UserName,
				ModelName:     name,
			}
			Global.Database.Create(&serverLog)
		}
		//log.Println(c.Request.RequestURI)
		//log.Println(c.Request.Body)
		//log.Println(c.Writer.Status())
		//log.Println(w.body.String())

	}
}
