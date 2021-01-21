package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @Summary 添加一篇文章
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

// @Summary 删除一篇文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Image not found"}"
// @Router /articles/{id} [DELETE]
// @Security ApiKeyAuth
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var article models.Article
	Global.Database.First(&article, id)
	if article.ID == 0 {
		c.JSON(404, gin.H{"message": "article not found"})
		return
	}
	Global.Database.Delete(&article)
	c.JSON(200, gin.H{"message": "delete success"})

}

// @Summary 更改一篇文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param   id      path int true  "文章id" default(1)
// @Param   article      body httpModels.NoIdArticle true  "历史事件"
// @Success 200 {object} models.Article
// @Failure 404 {string} string "{"message": "Article not found"}"
// @Router /articles/{id} [PUT]
// @Security ApiKeyAuth
func UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your article is not a number"})
		return
	}
	var article models.Article
	Global.Database.First(&article, id)
	if article.ID == 0 {
		c.JSON(404, gin.H{"message": "Article not found"})
		return
	}
	err = c.ShouldBindJSON(&article)
	if err != nil {
		log.Println(err)
	}
	if article.ID != id {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Save(&article)
	c.JSON(200, article)

}

// @Summary 获取所有文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Article
// @Router /articles [get]
func ReadArticles(c *gin.Context) {
	var articles []models.Article
	Global.Database.Find(&articles)
	c.JSON(200, articles)
}
