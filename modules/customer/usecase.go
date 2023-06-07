package customer

import (
	"database/sql"
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/customer"
	//do "github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer/do"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"github.com/imama2/bootcamp-bri-mini-project/utils/helper"
	"sync"
)

type UseCaseCustomerInterface interface {
	GetAllCustomer(dt do.Customer, pagi do.Pagination) (do.ListActorWithPaging, error)
	GetCustomerByID(dt do.Customer) (do.Customer, error)

	CreateCustomer(dt do.Customer) (int64, error)
	UpdateCustomerByID(dt do.Customer) (int64, error)
	DeleteCustomerByID(dt do.Customer) (int64, error)
}

type UseCaseCustomer struct {
	CustomerRepository repositories.CustomerRepositoryInterface
	DB                 *sql.DB
}

func NewCustomerUseCase(CustomerRepo repositories.CustomerRepositoryInterface, DB *sql.DB) UseCaseCustomerInterface {
	return &UseCaseCustomer{
		CustomerRepository: CustomerRepo,
		DB:                 DB,
	}
}

func (uc *UseCaseCustomer) CreateCustomer(dt do.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
		Avatar:    dt.Avatar,
	}
	i, err := uc.CustomerRepository.CreateCustomer(tx, et)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// DeleteCustomerByID implements CustomertUseCase.
func (uc *UseCaseCustomer) DeleteCustomerByID(dt do.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID: dt.ID,
	}
	i, err := uc.CustomerRepository.DeleteCustomerByID(tx, et)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetAllCustomer implements CustomertUseCase.
func (uc *UseCaseCustomer) GetAllCustomer(dt do.Customer, pagi do.Pagination) (do.ListActorWithPaging, error) {
	var err error
	var wg sync.WaitGroup
	var res []entity.Customer
	var resPaging entity.Pagination
	chRes := make(chan []entity.Customer, 1)
	chResPaging := make(chan entity.Pagination, 1)
	chErrRes := make(chan error, 1)
	chErrPaging := make(chan error, 1)

	// !Error tx with go routine, temporary solution using db queries
	// tx, err := uc.DB.Begin()
	// if err != nil {
	// 	return do.ListActorWithPaging{}, err
	// }
	// defer helper.CommitOrRollback(err, tx)

	// define pagination
	etPaging := entity.Pagination{
		Page:       pagi.Page,
		PerPage:    6,                   // always fix 6 data == LIMIT
		Total:      0,                   // after query
		TotalPages: 0,                   // after query, total / PerPage
		Offset:     (pagi.Page - 1) * 6, // (Page-1) * PerPage
	}

	et := entity.Customer{
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// get all customer with pagination
		res, err = uc.CustomerRepository.GetAllCustomer(uc.DB, et, etPaging)
		if err != nil {
			chErrRes <- err
		}
		chRes <- res
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Get Total Data
		resPaging, err = uc.CustomerRepository.Pagination(uc.DB, etPaging)
		if err != nil {
			chErrPaging <- err
		}
		chResPaging <- resPaging
	}()
	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case res = <-chRes:
			continue
		case resPaging = <-chResPaging:
			continue
		case err = <-chErrRes:
			return do.ListActorWithPaging{}, err
		case err = <-chErrPaging:
			return do.ListActorWithPaging{}, err
		}
	}

	totalPages := resPaging.Total / 6
	if resPaging.Total%6 != 0 {
		totalPages++
	}
	etPaging.Total = resPaging.Total
	etPaging.TotalPages = totalPages

	combineRes := do.ListActorWithPaging{
		Pagination: do.Pagination(etPaging),
		Customers:  DTOListCustomer(res),
	}

	return combineRes, nil
}

// GetCustomerByID implements CustomertUseCase.
func (uc *UseCaseCustomer) GetCustomerByID(dt do.Customer) (do.Customer, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return do.Customer{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID: dt.ID,
	}
	res, err := uc.CustomerRepository.GetCustomerByID(tx, et)
	if err != nil {
		return do.Customer{}, err
	}

	return DTOCustomer(res), nil
}

// UpdateCustomerByID implements CustomertUseCase.
func (uc *UseCaseCustomer) UpdateCustomerByID(dt do.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID:        dt.ID,
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
		Avatar:    dt.Avatar,
	}
	res, err := uc.CustomerRepository.UpdateCustomerByID(tx, et)
	if err != nil {
		return 0, err
	}

	return res, nil
}
