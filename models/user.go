package models

import "360-evaluation/utils"

type User struct {
	Model

	Name     string `json:"name" gorm:""`
	Password string `json:"Password" gorm:"Column=Password"`
}

func (u *User) SetPassword(password string) {
	u.Password = utils.Sha256([]byte(password))
}

func (u *User) ValidPassword(password string) bool {
	return u.Password == utils.Sha256([]byte(password))
}
