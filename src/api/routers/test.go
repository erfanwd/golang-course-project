package routers

import (
	"github.com/erfanwd/golang-course-project/api/handlers"
	"github.com/gin-gonic/gin"
)

func Test(r *gin.RouterGroup){
	handler := handlers.NewTestHandler()
	
	r.GET("/", handler.TestHandle)
	r.POST("/header-binder", handler.HeaderBinder1)
}