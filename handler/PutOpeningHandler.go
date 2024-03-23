package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
