package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 添加一个介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Param   passage      body httpModels.NoIdPassage true  "介绍"
// @Success 200 {object} models.Passage
// @Router /passage [POST]
// @Security ApiKeyAuth
func CreatePassage(c *gin.Context) {
	var passage models.Passage

	if err := c.BindJSON(&passage); err != nil {
		log.Println(err)
		c.JSON(400, "Not a Passage")
		return
	}

	if passage.ID != 0 {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}

	Global.Database.Create(&passage)
	c.JSON(200, passage)
}

// @Summary 获取介绍
// @Tags 介绍
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Passage
// @Router /passage [get]
func ReadPassage(c *gin.Context) {
	var passage models.Passage
	Global.Database.First(&passage)
	c.JSON(200, passage)
}
