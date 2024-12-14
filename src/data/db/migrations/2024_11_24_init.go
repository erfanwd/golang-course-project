package migrations

import (
	"errors"

	"github.com/erfanwd/golang-course-project/constants"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func Up1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, tables, country)
	tables = addNewTable(database, tables, city)
	tables = addNewTable(database, tables, user)
	tables = addNewTable(database, tables, role)
	tables = addNewTable(database, tables, userRole)

	database.Migrator().CreateTable(tables...)
	
}

func addNewTable(database *gorm.DB, tables []interface{}, table interface{}) []interface{} {
	if !database.Migrator().HasTable(table) {
		tables = append(tables, table)
	}
	createDefaultInformation(database)
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := &models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, adminRole)

	defaultRole := &models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExist(database, defaultRole)

	pass, _ := bcrypt.GenerateFromPassword([]byte(constants.DefaultAdminUserPassword), bcrypt.DefaultCost)
	defaultAdminUser := &models.User{
		Username: constants.DefaultAdminUsername,
		Password: string(pass),
	}
	createAdminUserIfNotExist(database, defaultAdminUser, adminRole.Id)
}

func createRoleIfNotExist(database *gorm.DB, role *models.Role) {
	var existingRole models.Role
	if err := database.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			database.Create(role)
		} 
	} else {
		*role = existingRole
	}
}

func createAdminUserIfNotExist(database *gorm.DB, user *models.User, roleId int) {
	var existingUser models.User
	if err := database.Where("username = ?", user.Username).First(&existingUser).Error; err != nil  {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			database.Create(user)
			ur := models.UserRole{UserId: user.Id, RoleId: roleId}
			database.Create(&ur)
		} 
	} else {
		*user = existingUser
	}
}
