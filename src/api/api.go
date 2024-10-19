package api

import (
	"fmt"

	"github.com/erfanwd/golang-course-project/api/middlewares"
	"github.com/erfanwd/golang-course-project/api/routers"
	"github.com/erfanwd/golang-course-project/api/validations"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitialServer(cfg *config.Config){
	r := gin.New()
	
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(),gin.Recovery())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%d",cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine){
	v1 := r.Group("/api/v1/")
	
	{
		healthCheck := v1.Group("health")
		routers.Health(healthCheck)

		test := v1.Group("test")
		routers.Test(test)
	}
}

func RegisterValidators(){
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile",validations.MobileNumberValidator,true)
		val.RegisterValidation("password",validations.PasswordValidator,true)
	}
}