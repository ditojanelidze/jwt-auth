package models

type User struct {
	//gorm.Model
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
