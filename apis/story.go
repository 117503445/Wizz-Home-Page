package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
)
// @Summary 获取所有历史事件
// @Tags stories
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Story
// @Router /stories [get]
func ReadStories(c *gin.Context) {
	var stories []models.Story
	Global.Database.Find(&stories)
	c.JSON(200, stories)
}
// @Summary 获取一个历史事件
// @Tags stories
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Success 200 {object} models.Story
// @Router /stories/{id} [get]
func ReadStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	Global.Database.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	}
	c.JSON(200, story)
}
// @Summary 添加一个历史事件
// @Tags stories
// @Accept  json
// @Produce  json
// @Param   story      body models.Story true  "历史事件"
// @Success 200 {object} models.Story
// @Router /stories [POST]
// @Security ApiKeyAuth
func CreateStory(c *gin.Context) {
	var story models.Story
	_ = c.BindJSON(&story)
	Global.Database.Create(&story)
	c.JSON(200, story)
}
// @Summary 更改一个历史事件
// @Tags stories
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Param   story      body models.Story true  "历史事件"
// @Success 200 {object} models.Story
// @Router /stories/{id} [PUT]
// @Security ApiKeyAuth
func UpdateStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	Global.Database.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	} else {
		_ = c.BindJSON(&story)
		Global.Database.Save(&story)
		c.JSON(200, story)
	}
}
// @Summary 删除一个历史事件
// @Tags stories
// @Accept  json
// @Produce  json
// @Param   id      path int true  "历史事件id" default(1)
// @Success 200 {object} models.Story
// @Router /stories/{id} [DELETE]
// @Security ApiKeyAuth
func DeleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story models.Story
	Global.Database.First(&story, id)
	if story.ID == 0 {
		c.JSON(404, gin.H{"message": "Story not found"})
		return
	} else {
		_ = c.BindJSON(&story)
		Global.Database.Delete(&story)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
