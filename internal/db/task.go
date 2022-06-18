package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	DueDate     time.Time `gorm:"index"`
	CompletedAt time.Time `gorm:"index"`
}
