package models

type Service struct {
	ID            int    `gorm:"primarikey" json:"id"`
	Technician_id int    `json:"technician_id"`
	Client_id     int    `json:"client_id"`
	Ac_id         int    `json:"ac_id"`
	Date          string `json:"date"`
	Status        string `json:"status"`
}
