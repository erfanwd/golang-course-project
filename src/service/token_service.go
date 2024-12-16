package service

import (
	"time"

	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/pkg/service_errors"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId    int
	FirstName string
	LastName  string
	Username  string
	email     string
	Roles     []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		logger: logging.NewLogger(cfg),
		cfg:    cfg,
	}
}


func (service *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireDuration = time.Now().Add(service.cfg.Jwt.AccessTokenExpireDuration*time.Second).Unix()
	td.RefreshTokenExpireDuration = time.Now().Add(service.cfg.Jwt.RefreshTokenExpireDuration*time.Second).Unix()

	atc := jwt.MapClaims{}

	atc["user_id"] = token.UserId
	atc["first_name"] = token.FirstName
	atc["last_name"] = token.LastName
	atc["username"] = token.Username
	atc["email"] = token.email
	atc["roles"] = token.Roles
	atc["exp"] = td.AccessTokenExpireDuration


	ac := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error

	td.AccessToken, err = ac.SignedString([]byte(service.cfg.Jwt.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}
	rtc["user_id"] = token.UserId
	rtc["exp"] = td.RefreshTokenExpireDuration


	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(service.cfg.Jwt.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil

}

func (service *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnexpectedError}
		}
		return []byte(service.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (service *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := service.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}