package article

import "leek-api/app/models"

type Content struct {
	models.BaseModel

	Markdown string `gorm:"type:longText;" json:"markdown"`
	Html     string `gorm:"type:longText;" json:"html"`

	ArticleID uint64 `gorm:"not null;index;" json:"article_id"`
}
