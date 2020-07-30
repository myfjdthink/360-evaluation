package models

import "fmt"

type User struct {
	Model

	Name string
}

func AddUser(name string) error {
	u := &User{
		Name: name,
	}

	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByName(name string) (*User, error) {
	var u User
	err := db.Select("Name").Where("Name = ?", name).First(&u).Error
	if err != nil {
		return nil, err
	}
	if u.Name == "" {
		err = fmt.Errorf("找不到用户:%s", name)
		return nil, err
	}
	return &u, nil
}
