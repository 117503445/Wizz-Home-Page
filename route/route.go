package route

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Middlewares"
	"Wizz-Home-Page/apis"
	_ "Wizz-Home-Page/docs" //引入 swagger 必备
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoute() {

	Global.Engine.Use(Middlewares.CORSMiddleware())

	logMiddleware := Middlewares.GetLogMiddlewareFunc()
	authMiddleware := Middlewares.GetAuthMiddleware()

	useHttps := viper.GetBool("useHttps")
	if useHttps {
		Global.Engine.Use(Middlewares.TlsHandler())
	}

	apiGroup := Global.Engine.Group("/api")

	auth := apiGroup.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)

	storyGroup := apiGroup.Group("/stories")
	//storyGroup.Use(logMiddleware)
	storyGroup.GET("", apis.ReadStories)
	storyGroup.GET("/:id", apis.ReadStory)

	storyGroup.POST("", authMiddleware.MiddlewareFunc(), logMiddleware, apis.CreateStory)
	storyGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.UpdateStory)
	storyGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.DeleteStory)

	PassageGroup := apiGroup.Group("/passage")
	PassageGroup.POST("", authMiddleware.MiddlewareFunc(), logMiddleware, apis.CreatePassage)

	productGroup := apiGroup.Group("/products")
	productGroup.GET("", apis.ReadProducts)
	productGroup.GET("/:id", apis.ReadProduct)
	productGroup.POST("", authMiddleware.MiddlewareFunc(), logMiddleware, apis.CreateProduct)
	productGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.UpdateProduct)
	productGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.DeleteProduct)

	memberGroup := apiGroup.Group("/members")
	memberGroup.GET("", apis.ReadMembers)
	memberGroup.GET("/:id", apis.ReadMember)
	memberGroup.POST("", authMiddleware.MiddlewareFunc(), logMiddleware, apis.CreateMember)
	memberGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.UpdateMember)
	memberGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), logMiddleware, apis.DeleteMember)

	imageGroup := apiGroup.Group("/image")
	imageGroup.GET("/UpToken", authMiddleware.MiddlewareFunc(), apis.GetUpToken)
	imageGroup.GET("/BackGroundImageUrls", apis.GetBackGroundImageUrls)
	imageGroup.GET("/PlaceAndDomain", apis.GetPlaceAndDomain)

	serverLogGroup := apiGroup.Group("/logs")
	serverLogGroup.GET("", apis.ReadServerLogs)

	backupGroup := apiGroup.Group("/backup")
	backupGroup.GET("", apis.ExportData)

	Global.Engine.GET("/ver", func(c *gin.Context) {
		c.JSON(200, "0205-1320")
	})

	Global.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
