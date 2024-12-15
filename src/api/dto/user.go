package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken                string `json:"accessToken`
	RefreshToken               string `json:"refreshToken"`
	AccessTokenExpireDuration  int64  `json:"accessTokenExpireDuration"`
	RefreshTokenExpireDuration int64  `json:"refreshTokenExpireDuration"`
}
