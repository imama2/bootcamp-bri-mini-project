package entity

type Approval struct {
	ID           int64
	AdminId      int64
	SuperAdminId int64
	Status       string
}
