package handlers

import (
	"net/http"

	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/api/helpers"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/service"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	CountryService *service.CountryService
}

func NewCountryHandler (cfg *config.Config) *CountryHandler {
	service := service.NewCountryService(cfg)
	return &CountryHandler{
		CountryService: service,
	}
}

// CreateCountry godoc
// @Summery create a country
// @Description create a country in db
// @Tags Country
// @Accept json
// @Produce json
// @Param Request body dto.CountryCreateOrUpdateRequest true "CountryCreateOrUpdateRequest"
// @Success 201 {object} helpers.BaseHttpResponse "Success"
// @Failure 400 {object} helpers.BaseHttpResponse "Failed"
// @Failure 409 {object} helpers.BaseHttpResponse "Failed"
// @Security BearerAuth
// @Router /v1/country/create [post]
func (h *CountryHandler) CreateCountry (ctx *gin.Context) {
	request := new(dto.CountryCreateOrUpdateRequest)
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseHttpResponseWithValidationError(nil, false, -1, err))
		return
	}
	response, err := h.CountryService.CreateCountry(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseHttpResponseWithError(nil, false, -1, err))
		return
	}
	ctx.JSON(http.StatusCreated, helpers.GenerateBaseHttpResponse(response, true, 0))
}