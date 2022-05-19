package model

import (
	"time"

	"gorm.io/gorm"
)

// Account represents main model of manipulation
type Account struct {
	ID        int32 `gorm:"primaryKey"`
	Owner     string
	Balance   float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
