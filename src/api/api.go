package api

import (
	"fmt"

	"github.com/erfanwd/golang-course-project/api/routers"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/gin-gonic/gin"
)

func InitialServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	v1 := r.Group("/api/v1/")
	
	{
		healthCheck := v1.Group("health")
		routers.Health(healthCheck)
	}

	r.Run(fmt.Sprintf(":%d",config.Server.Port))
}