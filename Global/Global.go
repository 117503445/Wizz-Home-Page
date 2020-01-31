package Global

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

var Engine *gin.Engine