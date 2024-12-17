package repo_interfaces

import (
	"context"

	"github.com/erfanwd/golang-course-project/data/models"
)

type UserRepositoryInterface interface {
	ExistsBy(attrName, attrValue string) (bool, error)
	GetDefaultRole() (int, error)
	Create(ctx context.Context, user *models.User) error
}
