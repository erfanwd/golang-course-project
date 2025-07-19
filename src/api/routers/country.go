package routers

import (
	"github.com/erfanwd/golang-course-project/api/handlers"
	"github.com/erfanwd/golang-course-project/api/middlewares"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/gin-gonic/gin"
)

func Country(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)
	router.POST("create", middlewares.Authentication(cfg), h.CreateCountry)
	router.GET("find/:countryId", middlewares.Authentication(cfg), h.GetById)
}
