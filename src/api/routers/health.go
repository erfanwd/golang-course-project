package routers

import (
	"github.com/erfanwd/golang-course-project/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()
	
	r.GET("/check", handler.HealthCheck)
}