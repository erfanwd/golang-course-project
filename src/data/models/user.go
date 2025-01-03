package models

type User struct {
	Id           int    `gorm:"primaryKey"`
	Username     string `gorm:"type:string;size:20;not null;unique"`
	FirstName    string `gorm:"type:string;size:15;null"`
	LastName     string `gorm:"type:string;size:25;null"`
	MobileNumber string `gorm:"type:string;size:11;null;unique;default:null"`
	Email        string `gorm:"type:string;size:64;null;unique;default:null"`
	Password     string `gorm:"type:string;size:64;not null"`
	Enabled      bool   `gorm:"default:true"`
	UserRoles    *[]UserRole
	BaseModel
}

type Role struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:string;size:10;not null,unique"`
	UserRoles *[]UserRole
	BaseModel
}

type UserRole struct {
	Id     int  `gorm:"primaryKey"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
	BaseModel
}
