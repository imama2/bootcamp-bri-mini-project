package entities

import (
	"gorm.io/gorm"
	"time"
)

// struktur data suatu table
type Customer struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Firstname string `gorm:"column:first_name"`
	Lastname  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Avatar    string `gorm:"column:avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
