package repositories

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type CustomerRepositoryInterface interface {
	GetByID(id int) []entities.Customer
}

func (repo CustomerRepository) GetByID(id int) []entities.Customer {
	// implementasi query get customer by id
	return []entities.Customer{}
}
