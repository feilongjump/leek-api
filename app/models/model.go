package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // json:"-" 在数据返回时，此类字段不返回
}
