package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var Ak string
var Sk string
var Bucket string

// @Summary 获取一个上传文件的upToken
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   fileName      query string true  "要上传的文件名,如 abc.png"
// @Success 200 {string} string "upToken"
// @Router /image/UpToken [get]
// @Security ApiKeyAuth
func GetUpToken(c *gin.Context) {
	keyToOverwrite := c.Query("fileName")
	//log.Println(keyToOverwrite)
	//log.Println(Ak)
	//log.Println(Sk)
	//log.Println(Bucket)
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", Bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket  )","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(Ak, Sk)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(200, upToken)
}

var BackGroundImageUrls []string

// @Summary 获取所有背景图片的url
// @Tags 图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string "["http://q52qkptnh.bkt.clouddn.com/1.png","http://q52qkptnh.bkt.clouddn.com/2.png","http://q52qkptnh.bkt.clouddn.com/3.png"]"
// @Router /image/BackGroundImageUrls [get]
func GetBackGroundImageUrls(c *gin.Context) {
	c.JSON(200, BackGroundImageUrls)
}
