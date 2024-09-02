package models

type AC struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Pk    string  `json:"pk"`
	Price float32 `json:"price"`
}
