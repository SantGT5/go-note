package handler

import (
	"net/http"

	"github.com/SantGT5/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Get openings
// @Description Get all openings
// @Tags Opening
// @Accept json
// @Produce json
// @Success 200 {object} GetOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "get-openings", openings)
}
