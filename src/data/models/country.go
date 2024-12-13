package models

type Country struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:15;type:string;not null"`
	Cities *[]City
	BaseModel
}
