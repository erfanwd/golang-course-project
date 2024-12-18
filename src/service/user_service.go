package service

import (
	"context"

	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/common"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/pkg/service_errors"
	"github.com/erfanwd/golang-course-project/repository"
	repo_interfaces "github.com/erfanwd/golang-course-project/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Logger       logging.Logger
	Cfg          *config.Config
	OtpService   *OtpService
	TokenService *TokenService
	UserRepo     repo_interfaces.UserRepositoryInterface
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{
		Logger:       logger,
		Cfg:          cfg,
		OtpService:   NewOtpService(cfg),
		TokenService: NewTokenService(cfg),
		UserRepo:     repository.NewUserRepo(cfg, logger),
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

func (service *UserService) RegisterByUsername(ctx context.Context, req *dto.RegisterUserByUsernameRequest) error {
	u := &models.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if err := service.checkUserExists("username", req.Username, service_errors.UsernameExists); err != nil {
		return err
	}

	if err := service.checkUserExists("email", req.Email, service_errors.EmailExists); err != nil {
		return err
	}

	hashedPassword, err := service.hashPassword(req.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return service.UserRepo.CreateUser(ctx, u)
}

func (service *UserService) LoginByUsername(ctx context.Context, req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	u := &models.User{
		Username: req.Username,
		Password: req.Password,
	}
	user, err := service.UserRepo.GetBy("username", req.Username)
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UsernameNotExists}
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(user.Password))
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.IncorrectPassword}
	}
	
	token, err := service.generateTokenForExistingUser(u.Username)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (service *UserService) RegisterLoginByMobileNumber(ctx context.Context, req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	u := &models.User{
		MobileNumber: req.MobileNumber,
		Username:     req.MobileNumber,
	}

	if err := service.OtpService.ValidateOtp(req.MobileNumber, req.Otp); err != nil {
		service.Logger.Error(logging.Otp, logging.OtpValidation, err.Error(), nil)
		return nil, err
	}

	exists, err := service.UserRepo.ExistsBy("mobile_number", req.MobileNumber)
	if err != nil {
		return nil, err
	}

	if exists {
		return service.generateTokenForExistingUser(u.Username)
	}

	hashedPassword, err := service.hashPassword(common.GeneratePassword())
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword

	if err := service.UserRepo.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return service.generateTokenForExistingUser(u.Username)
}

func (service *UserService) checkUserExists(field, value, errorMessage string) error {
	exists, err := service.UserRepo.ExistsBy(field, value)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: errorMessage}
	}
	return nil
}

func (service *UserService) hashPassword(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		service.Logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return "", err
	}
	return string(hp), nil
}

func (service *UserService) generateTokenForExistingUser(username string) (*dto.TokenDetail, error) {
	user, err := service.UserRepo.GetBy("username", username)
	if err != nil {
		return nil, err
	}

	tokenDto := &tokenDto{
		UserId:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Roles:     extractUserRoles(user.UserRoles),
	}

	token, err := service.TokenService.GenerateToken(tokenDto)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func extractUserRoles(userRoles *[]models.UserRole) []string {
	if userRoles == nil || len(*userRoles) == 0 {
		return nil
	}

	roles := make([]string, 0, len(*userRoles))
	for _, userRole := range *userRoles {
		roles = append(roles, userRole.Role.Name)
	}
	return roles
}
