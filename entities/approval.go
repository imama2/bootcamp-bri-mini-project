package entities

type approval struct {
	Id           uint   `gorm:"Primary_key"`
	SuperAdminId uint   `gorm:"column:super_admin_id"`
	Status       string `gorm:"column:status"`
}
