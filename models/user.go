package models

type User struct {
	Model

	Name string `json:"name" gorm:""`
	Age  int    `json:"age" gorm:""`
}
