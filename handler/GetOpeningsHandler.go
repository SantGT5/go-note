package handler

import (
	"net/http"

	"github.com/SantGT5/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "get-openings", openings)
}
