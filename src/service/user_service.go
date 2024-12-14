package service

import (
	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/common"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	Logger     logging.Logger
	Cfg        *config.Config
	Database   *gorm.DB
	OtpService *OtpService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		Database:   database,
		Logger:     logger,
		Cfg:        cfg,
		OtpService: NewOtpService(cfg),
	}
}


func (service *UserService) SendOtp(request *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := service.OtpService.SetOtp(request.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}