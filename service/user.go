package service

import (
	"360-evaluation/config"
	"360-evaluation/models"
	"errors"
)

func Hello(name string) string {
	return "hello " + name
}

func AddUser(name string) (err error, user models.User) {
	_, exist := getUserByName(name)
	if exist != nil {
		err = errors.New("用户名已存在")
		return
	}
	user.Name = name
	user.SetPassword(config.Config.Password)
	err = models.DB.Create(&user).Error
	return
}

func getUserByName(name string) (error, *models.User) {
	u := new(models.User)
	notRegister := models.DB.Where("name = ?", name).Take(&u).RecordNotFound()
	if notRegister {
		err := errors.New("找不到该用户名")
		return err, nil
	}
	return nil, u
}
