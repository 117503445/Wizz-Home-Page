package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summary 下载数据库
// @Tags 备份
// @Accept  json
// @Produce  json
// @Router /backup [get]
func ExportData(c *gin.Context) {
	filename := fmt.Sprintf("attachment; filename=%vWizz-Home-Page.Database", time.Now().Format("2006-01-02-15-04-05-"))
	fmt.Println(filename)
	c.Writer.Header().Add("Content-Disposition", filename)
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./data/Wizz-Home-Page.Database")
}
