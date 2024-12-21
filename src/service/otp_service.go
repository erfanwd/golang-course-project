package service

import (
	"fmt"
	"time"

	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/constants"
	"github.com/erfanwd/golang-course-project/data/cache"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/pkg/service_errors"
	"github.com/go-redis/redis/v7"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, redisClient: redis, cfg: cfg}
}

func (service *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[OtpDto](service.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(service.redisClient, key, val, service.cfg.Otp.Duration*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (service *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)

	res, err := cache.Get[OtpDto](service.redisClient, key)
	if err != nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpIsNotValid}
	} else if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpIsNotValid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(service.redisClient, key, res, service.cfg.Otp.Duration*time.Second)
		if err != nil {
			return err
		}
	}

	return nil

}
