package user

import (
	"leek-api/pkg/model"
)

func (u *User) Create() (err error) {

	if err = model.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func GetByUsername(username string) (user User, err error) {

	if err = model.DB.Where(User{Username: username}).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
