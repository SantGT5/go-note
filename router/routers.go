package router

import (
	"github.com/SantGT5/gopportunities/docs"
	"github.com/SantGT5/gopportunities/handler"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var basePath = "/api/v1"

func initializeRoutes(router *gin.Engine) {
	// init handler
	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.PutOpeningHandler)
		v1.GET("/openings", handler.GetOpeningsHandler)
	}

	// init swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
