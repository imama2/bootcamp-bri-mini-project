package entities

import "time"

type Account struct {
	ID         uint   `gorm:"primary_key"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	RoleId     uint   `gorm:"column:role_id"`
	Isverified bool   `gorm:"column:isverified"`
	Isactive   bool   `gorm:"column:isactive"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
