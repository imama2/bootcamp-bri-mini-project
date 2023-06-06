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
	UpdateCustomer(user *entities.Customer) (any, error)
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	DeleteCustomer(email string) (any, error)
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

func (repo Customer) UpdateCustomer(user *entities.Customer) (any, error) {
	err := repo.db.Model(&entities.Customer{}).
		Save(user).Error
	return nil, err
}

// DeleteCustomer by Id and email
func (repo Customer) DeleteCustomer(email string) (any, error) {
	err := repo.db.Model(&entities.Customer{}).
		Where("email = ?", email).
		Delete(&entities.Customer{}).
		Error
	return nil, err
}
