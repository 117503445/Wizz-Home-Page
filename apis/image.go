package apis

import
(
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/gin-gonic/gin"
	"fmt"
)



var Ak string
var Sk string
var Bucket string

// @Summary 获取一个上传文件的upToken
// @Tags 图片
// @Accept  json
// @Produce  json
// @Param   fileName      path string true  "要上传的文件名,如 abc.png"
// @Success 200 {string} string "upToken"
// @Router /image/UpToken [get]
func GetUpToken(c *gin.Context){
	keyToOverwrite := c.Params.ByName("fileName")
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", Bucket, keyToOverwrite),
	}
	mac := qbox.NewMac(Ak, Sk)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(200,upToken)
}
