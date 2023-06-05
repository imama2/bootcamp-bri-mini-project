package repositories

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

type CustomerRepositoryInterface interface {
	GetCustomerByID(id uint) (entities.Customer, error)
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
}

func (repo Customer) GetCustomerByID(id uint) (entities.Customer, error) {
	var customer entities.Customer
	repo.db.First(&customer, `id = ?`, id)
	return customer, nil
}

func (repo Customer) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	return customer, err
}
