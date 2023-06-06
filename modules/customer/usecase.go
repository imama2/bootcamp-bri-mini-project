package customer

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
)

type UseCaseCustomer interface {
	CreateCustomer(user CustomerParam) (entities.Customer, error)
	GetCustomerByID(id uint) (entities.Customer, error)
}

type useCaseCustomer struct {
	customerRepo repositories.CustomerRepositoryInterface
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) GetCustomerByID(id uint) (entities.Customer, error) {
	var cust entities.Customer
	cust, err := uc.customerRepo.GetCustomerByID(id)
	return cust, err
}
