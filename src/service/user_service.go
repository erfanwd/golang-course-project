package service

import (
	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/common"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/pkg/service_errors"
	"github.com/erfanwd/golang-course-project/repository"
	"golang.org/x/crypto/bcrypt"
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
		UserRepo:   repository.NewUserRepo(cfg, logger),
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

func (service *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := &models.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	exists, err := service.UserRepo.ExistsBy("username", req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	exists, err = service.UserRepo.ExistsBy("email", req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		service.Logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)
	
	service.UserRepo.Create(u)

	return nil
}
