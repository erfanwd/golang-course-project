package models

type City struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:15;type:string;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
	BaseModel
}
