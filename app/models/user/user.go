package user

import (
	"leek-api/app/models"
	"leek-api/pkg/password"
)

type User struct {
	models.BaseModel

	Username string `gorm:"type:varchar(255);not null;" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(255);not null;" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255);not null;" json:"-" binding:"required"`
}

// ComparePassword 对比密码匹配
func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}
