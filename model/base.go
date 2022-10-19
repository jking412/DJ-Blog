package model

import "time"

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
