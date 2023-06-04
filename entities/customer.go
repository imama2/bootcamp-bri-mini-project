package entities

import (
	"gorm.io/gorm"
	"time"
)

// struktur data suatu table
type Customer struct {
	gorm.Model
	id        uint   `gorm:"primary_key"`
	firstname string `gorm:"column:first_name"`
	lastname  string `gorm:"column:last_name"`
	email     string `gorm:"column:email"`
	avatar    string `gorm:"column:avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
