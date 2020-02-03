package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

// @Summary 获取一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Success 200 {object} models.Member
// @Failure 404 {string} string "{"message":"Member not found"}"
// @Router /members/{id} [get]
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

// @Summary 添加一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   member      body httpModels.NoIdMember true  "成员"
// @Success 200 {object} models.Member
// @Router /members [POST]
// @Security ApiKeyAuth
func CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.BindJSON(&member); err != nil {
		log.Println(err)
		c.JSON(400, "Not a Member")
		return
	}

	if member.ID != 0 {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Create(&member)
	c.JSON(200, member)
}

// @Summary 更改一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   member      body httpModels.NoIdMember true  "成员"
// @Success 200 {object} models.Member
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /members/{id} [PUT]
// @Security ApiKeyAuth
func UpdateMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var member models.Member
	Global.Database.First(&member, id)
	if member.ID == 0 {
		c.JSON(404, gin.H{"message": "Member not found"})
		return
	}
	err = c.ShouldBindJSON(&member)
	if err != nil {
		log.Println(err)
	}
	if member.ID != id {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Save(&member)
	c.JSON(200, member)

}

// @Summary 删除一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /members/{id} [DELETE]
// @Security ApiKeyAuth
func DeleteMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var member models.Member
	Global.Database.First(&member, id)
	if member.ID == 0 {
		c.JSON(404, gin.H{"message": "Member not found"})
		return
	}
	Global.Database.Delete(&member)
	c.JSON(200, gin.H{"message": "delete success"})

}
