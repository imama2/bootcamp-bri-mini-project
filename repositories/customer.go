package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/customer"
)

type CustomerRepository struct {
}

func NewCustomerRepository() CustomerRepositoryInterface {
	return &CustomerRepository{}
}

type CustomerRepositoryInterface interface {
	// pagination
	Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) // only Get Total Data
	GetAllCustomer(tx *sql.DB, et entity.Customer, etPaging entity.Pagination) ([]entity.Customer, error)
	// get count total data

	GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error)

	CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error)
	UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
	DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
}

func (repo *CustomerRepository) Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) {
	var res entity.Pagination

	SQL := `
	SELECT count(id) FROM
	customers c`
	varArgs := []interface{}{}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Pagination{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&res.Total)
		if err != nil {
			return entity.Pagination{}, err
		}
	}

	return res, nil
}

// CreateCustomer implements CustomerRepository.
func (repo *CustomerRepository) CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error) {
	SQL := `
	INSERT INTO customers(first_name, last_name, email, avatar) 
	VALUES (?, ?, ?, ?)`
	varArgs := []interface{}{
		et.FirstName,
		et.LastName,
		et.Email,
		et.Avatar,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return i, nil
}

// DeleteCustomerByID implements CustomerRepository.
func (repo *CustomerRepository) DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	SQL := `
	DELETE FROM
		customers
	WHERE
		id = ?`
	varArgs := []interface{}{
		et.ID,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetAllCustomer implements CustomerRepository.
func (repo *CustomerRepository) GetAllCustomer(tx *sql.DB, et entity.Customer, etPaging entity.Pagination) ([]entity.Customer, error) {
	result := make([]entity.Customer, 0)

	SQL := `
	SELECT id, first_name, last_name, email, avatar, created_at, updated_at
	FROM customers
	WHERE LOWER(first_name) LIKE ?
	OR LOWER(last_name) LIKE ?
	AND LOWER(email) like ?
	LIMIT ?, ?`
	varArgs := []interface{}{
		fmt.Sprintf("%%%s%%", et.FirstName),
		fmt.Sprintf("%%%s%%", et.LastName),
		fmt.Sprintf("%%%s%%", et.Email),
		etPaging.Offset,
		etPaging.PerPage,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.Customer
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Email, &res.Avatar, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// GetCustomerByID implements CustomerRepository.
func (repo *CustomerRepository) GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error) {
	res := entity.Customer{}

	SQL := `
	SELECT id, first_name, last_name, email, avatar, created_at, updated_at
	FROM customers
	WHERE id = ?`
	varArgs := []interface{}{
		et.ID,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Customer{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Email, &res.Avatar, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return entity.Customer{}, err
		}
	} else {
		return entity.Customer{}, errors.New("customer Not Found")
	}

	return res, nil
}

// UpdateCustomerByID implements CustomerRepository.
func (repo *CustomerRepository) UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	SQL := `
	UPDATE
		customers
	SET first_name=?, last_name=?, email=?, avatar=?
	WHERE
		id = ?`
	varArgs := []interface{}{
		et.FirstName,
		et.LastName,
		et.Email,
		et.Avatar,
		et.ID,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return i, nil
}
