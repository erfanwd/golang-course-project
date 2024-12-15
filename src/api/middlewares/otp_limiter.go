package middlewares

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)



func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(ctx *gin.Context) {
		limiter := limiter.GetLimiter(getIP(ctx.Request.RemoteAddr))
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, 
				helpers.GenerateBaseHttpResponseWithError(nil, false, -1, errors.New("not allowed")))
			ctx.Abort()	
		}else {
			ctx.Next()
		}
	}
}

func getIP(remoteAddr string) string {
    ip, _, err := net.SplitHostPort(remoteAddr)
    if err != nil {
        return remoteAddr
    }
    return ip
}