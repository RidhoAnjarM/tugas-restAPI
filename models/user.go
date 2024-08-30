package models

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Photo     string    `json:"photo"`
	Address   string    `json:"address"`
	Role      int       `json:"role"`
	
}

