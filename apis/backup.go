package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 下载数据库
// @Tags 备份
// @Accept  json
// @Produce  json
// @Router /backup [get]
func ExportData(c *gin.Context) {
	filename := fmt.Sprintf("attachment; filename=%s", "Wizz-Home-Page.Database")
	c.Writer.Header().Add("Content-Disposition", filename)
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./data/Wizz-Home-Page.Database")
}
