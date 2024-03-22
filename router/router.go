package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// init Router using the default Gin config
	router := gin.Default()

	// defining Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
