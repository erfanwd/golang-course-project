package repository

import (
	"context"

	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/constants"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	repo_interface "github.com/erfanwd/golang-course-project/repository/interfaces"
	"gorm.io/gorm"
)

var _ repo_interface.UserRepositoryInterface = &UserRepo{}

type UserRepo struct {
	*BaseRepo[models.User]
	Logger logging.Logger
}

func NewUserRepo(cfg *config.Config, logger logging.Logger) repo_interface.UserRepositoryInterface {
	database := db.GetDb()
	return &UserRepo{
		BaseRepo: NewBaseRepository[models.User](database),
		Logger:   logger,
	}
}

func (r *UserRepo) ExistsBy(attrName string, attrValue string) (bool, error) {
	var exists bool
	if err := r.Database.Model(&models.User{}).
		Select("count(*) > 0").
		Where(attrName+"= ?", attrValue).
		Find(&exists).Error; err != nil {
		r.Logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (r *UserRepo) GetBy(attrName string, attrValue string) (*models.User, error) {
	var user models.User
	if err := r.Database.Model(&models.User{}).
		Where(attrName+"= ?", attrValue).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).Find(&user).Error; err != nil {
		r.Logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetDefaultRole() (roleId int, err error) {
	if err := r.Database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		Find(&roleId).Error; err != nil {
		r.Logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return 0, err
	}
	return roleId, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.BaseRepo.WithTransaction(ctx, func(tx *gorm.DB) error {
		roleId, err := r.GetDefaultRole()
		if err != nil {
			return err
		}

		if err := tx.Create(user).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.UserRole{RoleId: roleId, UserId: user.Id}).Error; err != nil {
			return err
		}

		return nil
	})
}
