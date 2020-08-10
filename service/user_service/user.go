package user_service

import (
	"360-evaluation/models"
	"errors"
)

func Hello(name string) string {
	return "hello " + name
}

func AddUser(name string) (err error, user models.User) {
	notRegister := models.DB.Where("name = ?", name).First(&user).RecordNotFound()
	if !notRegister {
		err = errors.New("用户名已存在")
		return
	} else {
		user.Name = name
		err = models.DB.Create(&user).Error
	}
	return err, user
}
