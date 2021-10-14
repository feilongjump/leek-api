package article

import (
	"leek-api/app/models"
)

type Article struct {
	models.BaseModel

	Title  string `gorm:"type:varchar(255);not null;" json:"title" binding:"required"`
	UserID uint64 `gorm:"not null;index" json:"user_id"`

	Content Content
}
