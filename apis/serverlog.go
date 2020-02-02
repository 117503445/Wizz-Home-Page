package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"strconv"
)

// @Summary 分页获取日志
// @Tags 日志
// @Accept  json
// @Produce  json
// @Param   pageSize      query int true  "一页含有的log数量" default(5)
// @Param   pageIndex      query int true  "第几页" default(1)
// @Success 200 {string} string "{"LogTotalNum":10,"logs":[{"id":2,"TimeStamp":1580620678,"RequestMethod":"PUT","Username":"admin","RequestURI":"/api/products/1","ResponseCode":"200","ModelName":"一起来开黑"}],"pageTotalNum":10}"
// @Router /logs [get]
func ReadServerLogs(c *gin.Context) {
	pgIndex := c.Query("pageIndex")
	pageIndex, err := strconv.Atoi(pgIndex)
	if err != nil {
		log.Println(err)
		c.JSON(400, "pageIndex is not a number")
		return
	}

	pgSize := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pgSize)
	if err != nil {
		log.Println(err)
		c.JSON(400, "pageSize is not a number")
		return
	}

	var serverLogs []models.ServerLog

	Global.Database.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&serverLogs)
	var logNum int
	Global.Database.Model(&models.ServerLog{}).Count(&logNum)
	pageTotalNum := math.Ceil(float64(logNum) / float64(pageSize))
	c.JSON(200, gin.H{"logs": serverLogs, "LogTotalNum": logNum, "pageTotalNum": pageTotalNum})
}
