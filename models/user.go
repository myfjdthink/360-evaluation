package models

type User struct {
	Model

	Name string `json:"name" gorm:"comment:'用户名'"`
}
