package qiniu

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func GetUpToken(fileName string) string {
	// fileName -> fileName
	bucket := g.Cfg().GetString("qiniu.bucket")
	ak := g.Cfg().GetString("qiniu.ak")
	sk := g.Cfg().GetString("qiniu.sk")
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, fileName),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket  )","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
