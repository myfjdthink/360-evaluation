package service

import (
	"360-evaluation/utils"
	"errors"
)

func LoginByPassword(name string, password string) (err error, token string) {
	_, u := getUserByName(name)
	if u != nil && u.ValidPassword(password) {
		token, err = utils.GenerateToken(u.ID)
		return err, token
	}
	err = errors.New("错误的用户名或密码")
	return
}
