package service

import (
	"context"

	"github.com/erfanwd/golang-course-project/api/dto"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"github.com/erfanwd/golang-course-project/repository"
	repo_interfaces "github.com/erfanwd/golang-course-project/repository/interfaces"
)

type CountryService struct {
	Logger       logging.Logger
	Cfg          *config.Config
	CountryRepo  repo_interfaces.CountryRepositoryInterface
}

func NewCountryService(cfg *config.Config) *CountryService {
	logger := logging.NewLogger(cfg)
	return &CountryService{
		Logger:       logger,
		Cfg:          cfg,
		CountryRepo:  repository.NewCountryRepo(cfg, logger),
	}
}

func (s *CountryService) CreateCountry(ctx context.Context, req *dto.CountryCreateOrUpdateRequest) (*dto.CountryResponse, error) {
	entity := &models.Country{
		Name: req.Name,
	} 
	
	country, err := s.CountryRepo.Create(ctx, entity)

	if err != nil {
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	res := &dto.CountryResponse{
		Id: country.Id,
		Name: country.Name,
	}
	return res, nil
}

func (s * CountryService) GetById(ctx context.Context, countryId int) (*dto.CountryResponse, error) {
	country, err := s.CountryRepo.FindByID(ctx, uint(countryId))
	if err != nil {
		s.Logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	response := &dto.CountryResponse{
		Id: country.Id,
		Name: country.Name,
	}
	return response, nil 
}
