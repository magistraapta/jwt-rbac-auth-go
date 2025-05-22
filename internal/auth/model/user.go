package model

type User struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"roles" gorm:"not null"`
}
