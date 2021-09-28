package models

import "time"

type BaseModel struct {
	ID        uint64
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	DeletedAt time.Time `gorm:"index"`
}
