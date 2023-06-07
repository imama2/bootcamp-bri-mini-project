package entity

import "time"

type Account struct {
	ID         int64
	Username   string
	Password   string
	RoleId     int64
	IsVerified bool
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
