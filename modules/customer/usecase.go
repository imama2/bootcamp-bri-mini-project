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

func (uc useCaseCustomer) UpdateCustomer(param CustomerParam, id uint) (any, error) {
	var editCustomer *entities.Customer
	editCustomer = &entities.Customer{
		ID:        id,
		Firstname: param.Firstname,
		Lastname:  param.Lastname,
		Email:     param.Email,
		Avatar:    param.Avatar,
	}
	_, err := uc.customerRepo.UpdateCustomer(editCustomer)
	if err != nil {
		return *editCustomer, err
	}
	return *editCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomer(email string) (any, error) {
	_, err := uc.customerRepo.DeleteCustomer(email)
	return nil, err
}
