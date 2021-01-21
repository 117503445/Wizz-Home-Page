package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 添加一个文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   article      body httpModels.NoIdArticle true  "文章"
// @Success 200 {object} models.Article
// @Router /articles [POST]
// @Security ApiKeyAuth
func CreateArticle(c *gin.Context) {
	var article models.Article

	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.JSON(400, "Not a Article")
		return
	}

	if article.ID != 0 {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Create(&article)
	c.JSON(200, article)
}
