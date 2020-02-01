package Middlewares

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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
		//w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		//c.Writer = w

		timestamp := time.Now().Unix()
		method := c.Request.Method
		name := "Can't find name"
		uri := c.Request.RequestURI
		if method == "PUT" || method == "DELETE" {
			id := c.Params.ByName("id")
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
			var b []byte
			b, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println(err)
			}

			err=c.Request.Body.Close()
			if err != nil {
				log.Println(err)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

			err = json.Unmarshal(b, &nm)
			if err != nil {
				log.Println(err)
			}
			name = nm.Name
		}

		c.Next()

		user, exists := c.Get(identityKey)
		if exists {
			code := strconv.Itoa(c.Writer.Status())
			serverLog := models.ServerLog{
				TimeStamp:     timestamp,
				RequestMethod: c.Request.Method,
				RequestURI:    uri,
				ResponseCode:  code,
				Username:      user.(*User).UserName,
				ModelName:     name,
			}

			Global.Database.Create(&serverLog)
		}
	}
}
