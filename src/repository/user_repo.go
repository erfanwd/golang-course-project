package repository

import (
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/models"
	"github.com/erfanwd/golang-course-project/pkg/logging"
	"gorm.io/gorm"
)

type UserRepo struct {
	Logger    logging.Logger
	Database  *gorm.DB
	UserModel *models.User
}

func NewUserRepo(cfg *config.Config, logger logging.Logger) *UserRepo {
	database := db.GetDb()
	return &UserRepo{
		Database:  database,
		Logger:    logger,
		UserModel: &models.User{},
	}
}

func (r *UserRepo) ExistsByEmail(email string) (bool, error) {
	var exists bool
	if err := r.Database.Model(r.UserModel).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error; err != nil {
			r.Logger.Error(logging.Postgres, logging.Select, err.Error(),nil)
			return false, err
	}
	return exists, nil 
}
