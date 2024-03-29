package router

import (
	"github.com/SantGT5/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// init handler
	handler.InitializeHandler()

	var v1 *gin.RouterGroup = router.Group("/api/v1/")

	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.PutOpeningHandler)
		v1.GET("/openings", handler.GetOpeningsHandler)
	}
}
