package models


type Role struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}