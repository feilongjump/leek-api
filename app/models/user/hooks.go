package user

import (
	"gorm.io/gorm"
	"leek-api/pkg/password"
)

// BeforeSave GORM 的模型钩子，更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	u.Password = password.Hash(u.Password)

	return
}
