package repo_interfaces

import (
	"context"

	"github.com/erfanwd/golang-course-project/data/models"
)

type UserRepositoryInterface interface {
	ExistsBy(attrName string, attrValue string) (bool, error)
	GetDefaultRole() (int, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetBy(attrName string, attrValue string) (*models.User, error)
}
