package models

type User struct {
	Id       int64
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
