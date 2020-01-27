package apis

import (
	"Wizz-homepage-go/Global"
	"Wizz-homepage-go/models"
	"github.com/gin-gonic/gin"
)

func ReadProducts(c *gin.Context){
	var products []models.Product
	Global.Database.Find(&products)
	c.JSON(200,products)
}

func ReadProduct(c *gin.Context){
	id := c.Params.ByName("id")
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	_ = c.BindJSON(&product)
	Global.Database.Create(&product)
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	} else {
		_ = c.BindJSON(&product)
		Global.Database.Save(&product)
		c.JSON(200, product)
	}
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	Global.Database.First(&product, id)
	if product.ID == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	} else {
		_ = c.BindJSON(&product)
		Global.Database.Delete(&product)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}
