package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
)
// @Summary 获取所有成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Member
// @Router /members [get]
func ReadMembers(c *gin.Context) {
	var members []models.Member
	Global.Database.Find(&members)
	c.JSON(200, members)
}

func ReadMember(c *gin.Context) {
	id := c.Params.ByName("id")
	var member models.Member
	Global.Database.First(&member, id)
	if member.ID == 0 {
		c.JSON(404, gin.H{"message": "Member not found"})
		return
	}
	c.JSON(200, member)
}

func CreateMember(c *gin.Context) {
	var member models.Member
	_ = c.BindJSON(&member)
	Global.Database.Create(&member)
	c.JSON(200, member)
}

func UpdateMember(c *gin.Context) {
	id := c.Params.ByName("id")
	var member models.Member
	Global.Database.First(&member, id)
	if member.ID == 0 {
		c.JSON(404, gin.H{"message": "Member not found"})
		return
	} else {
		_ = c.BindJSON(&member)
		Global.Database.Save(&member)
		c.JSON(200, member)
	}
}

func DeleteMember(c *gin.Context) {
	id := c.Params.ByName("id")
	var member models.Member
	Global.Database.First(&member, id)
	if member.ID == 0 {
		c.JSON(404, gin.H{"message": "Member not found"})
		return
	} else {
		Global.Database.Delete(&member)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
