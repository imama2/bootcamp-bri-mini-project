package entities

type roles struct {
	Id       uint   `gorm:"Primary_key"`
	RoleName string `gorm:"column:role_name"`
}
