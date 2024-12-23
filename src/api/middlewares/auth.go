package middlewares

import (
	"net/http"
	"strings"

	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/constants"
	"github.com/erfanwd/golang-course-project/pkg/service_errors"
	"github.com/erfanwd/golang-course-project/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = service.NewTokenService(cfg)
	return func(ctx *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := ctx.GetHeader("Authorization")
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.GenerateBaseHttpResponseWithError(nil, false, -1, err))
			return
		}
		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		ctx.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		ctx.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		ctx.Set(constants.EmailKey, claimMap[constants.EmailKey])
		ctx.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])
		ctx.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		ctx.Next()

	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				helpers.GenerateBaseHttpResponse(nil, false, -1))
			return
		}
		valRoles := ctx.Keys[constants.RolesKey]
		if valRoles == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				helpers.GenerateBaseHttpResponse(nil, false, -1))
			return
		}
		val := valRoles.([]interface{})
		finalItems := map[string]int{}
		for _, role := range val {
			finalItems[role.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := finalItems[item]; ok {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden,
			helpers.GenerateBaseHttpResponse(nil, false, -1))
	}
}
