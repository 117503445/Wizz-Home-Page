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
// @Router /members/update/{id} [PUT]
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

// @Summary 上移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} models.Member
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /members/up/{id} [PUT]
// @Security ApiKeyAuth
func UpMember(c *gin.Context) {
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

	membertype, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your type is not a number"})
		return
	} else if membertype <= 0 && membertype > 4 {
		c.JSON(400, gin.H{"message": "your type is not correct"})
		return
	} else if member.MemberType != membertype {
		c.JSON(400, gin.H{"message": "your type do not match your id"})
		return
	}
	var members []models.Member
	Global.Database.Where("member_type = ?", membertype).Find(&members)
	if len(members) == 0 || len(members) == 1 {
		c.JSON(400, gin.H{"message": " the number of member is 0 or 1"})
		return
	} else if members[0].ID == id {
		c.JSON(400, gin.H{"message": " the first member can not up"})
		return
	}
	i := 1
	var m1 models.Member
	var m2 models.Member
	m1 = members[i]
	for m1.ID != id {
		i++
		m1 = members[i]
	}
	m2 = members[i-1]
	m1, m2 = m2, m1
	m1.ID, m2.ID = m2.ID, m1.ID
	Global.Database.Save(&m1)
	Global.Database.Save(&m2)
	Global.Database.Where("member_type = ?", membertype).Find(&members)
	c.JSON(200, members)

}

// @Summary 下移一个成员
// @Tags 成员
// @Accept  json
// @Produce  json
// @Param   id      path int true  "成员id" default(1)
// @Param   type	query	int true  "成员类型"
// @Success 200 {array} models.Member
// @Failure 404 {string} string "{"message": "Member not found"}"
// @Router /members/down/{id} [PUT]
// @Security ApiKeyAuth
func DownMember(c *gin.Context) {
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

	membertype, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your type is not a number"})
		return
	} else if membertype <= 0 && membertype > 4 {
		c.JSON(400, gin.H{"message": "your type is not correct"})
		return
	} else if member.MemberType != membertype {
		c.JSON(400, gin.H{"message": "your type do not match your id"})
		return
	}
	var members []models.Member
	Global.Database.Where("member_type = ?", membertype).Find(&members)
	if len(members) == 0 || len(members) == 1 {
		c.JSON(400, gin.H{"message": " the number of member is 0 or 1"})
		return
	} else if members[len(members)-1].ID == id {
		c.JSON(400, gin.H{"message": " the last member can not up"})
		return
	}
	i := 0
	var m1 models.Member
	var m2 models.Member
	m1 = members[i]
	for m1.ID != id {
		i++
		m1 = members[i]
	}
	m2 = members[i+1]
	m1, m2 = m2, m1
	m1.ID, m2.ID = m2.ID, m1.ID
	Global.Database.Save(&m1)
	Global.Database.Save(&m2)
	Global.Database.Where("member_type = ?", membertype).Find(&members)
	c.JSON(200, members)

}
