package middlewares

import (
	"net/http"

	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helpers.GenerateBaseHttpResponseWithError(nil, false, -1, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helpers.GenerateBaseResponseWithAnyError(nil, false, -1, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}