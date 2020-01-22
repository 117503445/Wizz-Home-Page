package story

import (
	"Wizz-homepage-go/Global"
	"Wizz-homepage-go/models"
	"github.com/gin-gonic/gin"
)

func ReadStories(c *gin.Context) {

	var stories []models.Story
	Global.Database.Find(&stories)
	c.JSON(200, stories)
}

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

func CreateStory(c *gin.Context) {
	var story models.Story
	_ = c.BindJSON(&story) //绑定一个请求主体到一个类型
	Global.Database.Create(&story)
	c.JSON(200, story)
}

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
