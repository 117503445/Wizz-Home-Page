package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
)
// @Summary 获取所有日志
// @Tags 日志
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ServerLog
// @Router /logs [get]
func ReadServerLogs(c *gin.Context){
	var serverLogs []models.ServerLog
	Global.Database.Find(&serverLogs)
	c.JSON(200, serverLogs)
}