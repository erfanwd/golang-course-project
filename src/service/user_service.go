package service

import (
	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/common"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/repository"
)

type UserService struct {
	Logger     logging.Logger
	Cfg        *config.Config
	OtpService *OtpService
	UserRepo   *repository.UserRepo
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{
		Logger:     logger,
		Cfg:        cfg,
		OtpService: NewOtpService(cfg),
		UserRepo: repository.NewUserRepo(cfg, logger),
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
