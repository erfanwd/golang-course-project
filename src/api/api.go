package api

import (
	"fmt"

	"github.com/erfanwd/golang-course-project/api/routers"
	"github.com/erfanwd/golang-course-project/api/validations"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitialServer(){
	config := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile",validations.MobileNumberValidator,true)
	}
	r.Use(gin.Logger(),gin.Recovery())

	v1 := r.Group("/api/v1/")
	
	{
		healthCheck := v1.Group("health")
		routers.Health(healthCheck)

		test := v1.Group("test")
		routers.Test(test)
	}

	r.Run(fmt.Sprintf(":%d",config.Server.Port))
}