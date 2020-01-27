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