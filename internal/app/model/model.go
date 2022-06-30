package model

import (
	"gorm.io/gorm"
	"time"
)

func Tables() []interface{} {
	return []interface{}{}
}

type Model struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated" gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
