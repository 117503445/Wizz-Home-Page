package Middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 封装 gin.ResponseWriter, gin.Context 进行 Write 时也对 body 进行了 Write 操作
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
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

		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Println(c.ClientIP())
		log.Println(latency)
		log.Println(c.Request.Method)
		log.Println(c.Request.Header)
		user, exists := c.Get(identityKey)
		if exists{
			log.Println(user.(*User).UserName)
		}
		log.Println(c.Request.RequestURI)
		log.Println(c.Request.Body)
		log.Println(c.Writer.Status())
		log.Println(w.body.String()) //Response body
	}
}
