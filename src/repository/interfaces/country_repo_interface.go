package repo_interfaces

import "github.com/erfanwd/golang-course-project/data/models"

type CountryRepositoryInterface interface {
	BaseRepositoryInterface[models.Country]
}
