package repository

import (
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	repo_interface "github.com/erfanwd/golang-course-project/repository/interfaces"
)

var _ repo_interface.CountryRepositoryInterface = &CountryRepo{}

type CountryRepo struct {
	*BaseRepo[models.Country]
	Logger logging.Logger
}

func NewCountryRepo(cfg *config.Config, logger logging.Logger) repo_interface.CountryRepositoryInterface {
	database := db.GetDb()
	return &CountryRepo{
		BaseRepo: NewBaseRepository[models.Country](database),
		Logger:   logger,
	}
}
