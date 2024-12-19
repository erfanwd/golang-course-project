package routers

import (
	"github.com/erfanwd/golang-course-project/api/handlers"
	"github.com/erfanwd/golang-course-project/api/middlewares"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("login-by-username", h.LoginByUsername)
}
