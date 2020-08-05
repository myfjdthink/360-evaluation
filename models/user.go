package models

import "fmt"

type User struct {
	Model

	Name string
}

func AddUser(name string) (*User, error) {
	u := &User{
		Name: name,
	}

	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func GetUserByName(name string) (*User, error) {
	var u User
	err := db.Select("name").Where("name = ?", name).First(&u).Error
	if err != nil {
		return nil, err
	}
	if u.Name == "" {
		err = fmt.Errorf("找不到用户:%s", name)
		return nil, err
	}
	return &u, nil
}
