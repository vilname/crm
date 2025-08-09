package config

import (
	"api/src/controller/rest"
	"api/src/util/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(middleware.EnableCORS)
	//router.Use(middleware.Auth)

	router.POST("/pdf/text", rest.TextFromPdf)

	answer := router.Group("/answer")
	answer.GET("/list", rest.ListAnswer)
	answer.GET("/get/:id", rest.GetAnswer)
	answer.POST("/create", rest.CreateAnswer)

// 	docs.SwaggerInfo.BasePath = ""
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
