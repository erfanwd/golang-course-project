package handlers

import (
	"net/http"

	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/service"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *service.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := service.NewUserService(cfg)
	return &UsersHandler{
		service: service,
	}
}

// SendOtp godoc
// @Summery send otp to user
// @Description send otp to user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helpers.BaseHttpResponse "Success"
// @Failure 400 {object} helpers.BaseHttpResponse "Failed"
// @Failure 409 {object} helpers.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (handler *UsersHandler) SendOtp(c *gin.Context) {
	request := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseHttpResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = handler.service.SendOtp(request)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseHttpResponseWithError(nil, false, -1, err))
		return
	}
	// Call Sms Service
	c.JSON(http.StatusCreated, helpers.GenerateBaseHttpResponse(nil, true, 0))
}
