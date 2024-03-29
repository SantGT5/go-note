package handler

import (
	"fmt"
	"net/http"

	"github.com/SantGT5/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {

	body := gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
	}

	if data != nil {
		body["data"] = data
	}

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, body)
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type GetOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type GetOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string `json:"message"`
}
