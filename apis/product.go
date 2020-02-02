package apis

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @Summary 获取所有产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Router /products [get]
func ReadProducts(c *gin.Context) {
	var products []models.Product
	Global.Database.Find(&products)
	c.JSON(200, products)
}

// @Summary 获取一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Success 200 {object} models.Product
// @Failure 404 {string} string "{"message":"Product not found"}"
// @Router /products/{id} [get]
func ReadProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, product)
}

// @Summary 添加一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   product      body httpModels.NoIdProduct true  "产品"
// @Success 200 {object} models.Product
// @Router /products [POST]
// @Security ApiKeyAuth
func CreateProduct(c *gin.Context) {
	var product models.Product
	_ = c.BindJSON(&product)
	if product.ID != 0 {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Create(&product)
	c.JSON(200, product)
}

// @Summary 更改一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Param   product      body httpModels.NoIdProduct true  "产品"
// @Success 200 {object} models.Product
// @Failure 404 {string} string "{"message": "Product not found"}"
// @Router /products/{id} [PUT]
// @Security ApiKeyAuth
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	err = c.ShouldBindJSON(&product)
	if err != nil {
		log.Println(err)
	}
	if product.ID != id {
		c.JSON(400, gin.H{"message": "Pass id in body is not allowed"})
		return
	}
	Global.Database.Save(&product)
	c.JSON(200, product)

}

// @Summary 删除一个产品
// @Tags 产品
// @Accept  json
// @Produce  json
// @Param   id      path int true  "产品id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Product not found"}"
// @Router /products/{id} [DELETE]
// @Security ApiKeyAuth
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "your id is not a number"})
		return
	}
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	Global.Database.Delete(&product)
	c.JSON(200, gin.H{"message": "delete success"})

}
