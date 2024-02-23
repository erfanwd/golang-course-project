package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck (context *gin.Context){
	context.JSON(200,"healthy")
}