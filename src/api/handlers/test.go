package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {}
type Body struct {
	Id int `json:"id" binding:"required"` 
	Name string `json:"name" binding:"required"` 
	Mobile string `json:"mobile" binding:"mobile,required"` 
}
func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) TestHandle (context *gin.Context){
	context.JSON(200,gin.H{
		"result":"test result",
	})
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context){
	body := Body{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"result":err.Error(),
		})
		return
	}
	ctx.JSON(200,gin.H{
		"result":"test result",
		"header":body,
	})
}