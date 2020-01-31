package Middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetLogMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Using mid logger")
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Println(latency)
		log.Println(c.Writer.Status())
	}
}
