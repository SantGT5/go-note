package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// init Router
	var router *gin.Engine = gin.Default()

	// init Routers
	initializeRoutes(router)

	// listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
