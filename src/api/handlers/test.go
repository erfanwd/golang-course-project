package handlers

import (
	"net/http"

	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/gin-gonic/gin"
)

type TestHandler struct{}
type Body struct {
	Id     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required,min=4"`
	Mobile string `json:"mobile" binding:"mobile,required"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) TestHandle(context *gin.Context) {
	context.JSON(200, gin.H{
		"result": "test result",
	})
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	body := Body{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseHttpResponseWithValidationError("", false, 0, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GenerateBaseHttpResponse(body, true, 0))
}
