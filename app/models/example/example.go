package example

import "leek-api/app/models"

type Example struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;"`
	Body  string `gorm:"type:varchar(255);"`
}
