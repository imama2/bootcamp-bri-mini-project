package domain

type Approval struct {
	ID           int64
	AdminId      int64
	SuperAdminID int64
	Status       string
}