package repository

import (
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/constants"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"gorm.io/gorm"
)

type UserRepo struct {
	Logger    logging.Logger
	Database  *gorm.DB
}

func NewUserRepo(cfg *config.Config, logger logging.Logger) *UserRepo {
	database := db.GetDb()
	return &UserRepo{
		Database:  database,
		Logger:    logger,
	}
}

func (r *UserRepo) ExistsBy(attrName string, attrValue string) (bool, error) {
	var exists bool
	if err := r.Database.Model(&models.User{}).
		Select("count(*) > 0").
		Where(attrName + "= ?", attrValue).
		Find(&exists).Error; err != nil {
			r.Logger.Error(logging.Postgres, logging.Select, err.Error(),nil)
			return false, err
	}
	return exists, nil 
}

func (r *UserRepo) GetDefaultRole() (roleId int, err error) {
	if err := r.Database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		Find(&roleId).Error; err != nil {
			r.Logger.Error(logging.Postgres, logging.Select, err.Error(),nil)
			return 0, err
	}
	return roleId, nil 
}

func (r *UserRepo) Create(user *models.User) error {
	roleId, err := r.GetDefaultRole()
	if err != nil {
		r.Logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}
	tx := r.Database.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		r.Logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}

	if err := tx.Create(&models.UserRole{RoleId: roleId, UserId: user.Id}).Error; err != nil {
		tx.Rollback()
		r.Logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}

	tx.Commit()
	return nil

}
