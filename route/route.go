package route

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Middlewares"
	"Wizz-Home-Page/apis"
	_ "Wizz-Home-Page/docs" //引入 swagger 必备
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoute() {
	logMiddleware := Middlewares.GetLogMiddlewareFunc()
	authMiddleware := Middlewares.GetAuthMiddleware()

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

	productGroup := apiGroup.Group("/products")
	productGroup.GET("", apis.ReadProducts)
	productGroup.GET("/:id", apis.ReadProduct)
	productGroup.POST("", authMiddleware.MiddlewareFunc(), apis.CreateProduct)
	productGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), apis.UpdateProduct)
	productGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), apis.DeleteProduct)

	memberGroup := apiGroup.Group("/members")
	memberGroup.GET("", apis.ReadMembers)
	memberGroup.GET("/:id", apis.ReadMember)
	memberGroup.POST("", authMiddleware.MiddlewareFunc(), apis.CreateMember)
	memberGroup.PUT("/:id", authMiddleware.MiddlewareFunc(), apis.UpdateMember)
	memberGroup.DELETE("/:id", authMiddleware.MiddlewareFunc(), apis.DeleteMember)

	serverLogGroup := apiGroup.Group("/logs")
	serverLogGroup.GET("", apis.ReadServerLogs)

	Global.Engine.GET("/ver", func(c *gin.Context) {
		c.JSON(200, "0131-1253")
	})

	Global.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
