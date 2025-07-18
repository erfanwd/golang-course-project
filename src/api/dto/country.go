package dto

type CountryCreateOrUpdateRequest struct {
	Name string `json:"name" binding:"required,unique=countries,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

