package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	//"github.com/imama2/bootcamp-bri-mini-project/entities/account"
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/account"
	//do "github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"gorm.io/gorm"
)

type AccountRepository struct{}

func NewAccountRepository() AccountRepositoryInterface {
	return &AccountRepository{}
}

func (repo *AccountRepository) Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) {
	var res entity.Pagination
	return res, nil
}

type Account struct {
	db *gorm.DB
}

func NewAccount(dbCrud *gorm.DB) Account {
	return Account{
		db: dbCrud,
	}
}

type AccountRepositoryInterface interface {
	AccountAuthentication(tx *sql.Tx, account entity.Account) (entity.Account, error)
	SignIn(tx *sql.Tx, token entity.Token) (string, error) // TODO: store token

	// account
	AccountRegistration(tx *sql.Tx, account entity.Account) (int64, error)
	GetAllAdmin(tx *sql.DB, account entity.Account, et entity.Pagination) ([]entity.Account, error)
	Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) // only Get Total Data

	// admin reg
	ApprovalAdd(tx *sql.Tx, addApp entity.Approval) (int64, error)

	// super_admin only
	GetAllApprovalAdmin(tx *sql.Tx) ([]entity.Approval, error)
	UpdateAdminRegStatusByAdminID(tx *sql.Tx, addApp entity.Approval) (int64, error)
	UpdateAdminStatusByAdminID(tx *sql.Tx, account entity.Account) (int64, error)
	DeleteAdminByID(tx *sql.Tx, account entity.Account) (int64, error)
}

// GetAllAdmin implements AccountRepository.
func (repo *AccountRepository) GetAllAdmin(tx *sql.DB, actor entity.Account, etPage entity.Pagination) ([]entity.Account, error) {
	result := make([]entity.Account, 0)

	SQL := `
	SELECT id, username, role_id, is_verified, is_active, created_at, updated_at
	FROM actors
	WHERE LOWER(username) LIKE ?
	AND role_id = 1
	LIMIT ?, ?`
	varArgs := []interface{}{
		fmt.Sprintf("%%%s%%", actor.Username),
		etPage.Offset,
		etPage.PerPage,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.Account
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.Username, &res.RoleId, &res.IsVerified, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// DeleteAdminByID implements AccountRepository.
func (repo *AccountRepository) DeleteAdminByID(tx *sql.Tx, actor entity.Account) (int64, error) {
	SQL := `
	DELETE FROM
		actors
	WHERE
		id = ?`
	varArgs := []interface{}{
		actor.ID,
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

// UpdateAdminStatusByAdminID implements AccountRepository.
func (repo *AccountRepository) UpdateAdminStatusByAdminID(tx *sql.Tx, actor entity.Account) (int64, error) {
	SQL := `
	UPDATE actors 
	SET is_verified=?, is_active=? 
	WHERE id = ?`
	varArgs := []interface{}{
		actor.IsVerified,
		actor.IsActive,
		actor.ID,
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

// UpdateAdminRegStatusByAdminID implements AccountRepository.
func (repo *AccountRepository) UpdateAdminRegStatusByAdminID(tx *sql.Tx, adminReg entity.Approval) (int64, error) {
	SQL := `
	UPDATE admin_reg 
	SET status=?
	WHERE admin_id = ?`
	varArgs := []interface{}{
		adminReg.Status,
		adminReg.AdminId,
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

// GetAllApprovalAdmin implements AccountRepository.
func (repo *AccountRepository) GetAllApprovalAdmin(tx *sql.Tx) ([]entity.Approval, error) {
	result := make([]entity.Approval, 0)

	SQL := `
	SELECT id, admin_id, super_admin_id, status
	FROM admin_reg `

	rows, err := tx.Query(SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.Approval
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.AdminId, &res.SuperAdminId, &res.Status)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// VerifyActorCredential implements AccountRepository.
func (repo *AccountRepository) AccountAuthentication(tx *sql.Tx, actor entity.Account) (entity.Account, error) {
	SQL := `
	SELECT id, username, password, role_id, is_verified, is_active, created_at, updated_at
	FROM actors 
	WHERE username = ?`
	varArgs := []interface{}{
		actor.Username,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Account{}, err
	}
	defer rows.Close()

	res := entity.Account{}
	if rows.Next() {
		err := rows.Scan(&res.ID, &res.Username, &res.Password, &res.RoleId, &res.IsVerified, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return entity.Account{}, err
		}
	} else {
		return entity.Account{}, errors.New("incorrect username or password")
	}

	return res, nil
}

// RegisterAdmin implements AccountRepository.
func (repo *AccountRepository) ApprovalAdd(tx *sql.Tx, adminReg entity.Approval) (int64, error) {
	SQL := `
	INSERT INTO admin_reg(admin_id, super_admin_id, status) 
	VALUES (?, 1, ?)`
	varArgs := []interface{}{
		adminReg.AdminId,
		adminReg.Status,
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

// AddActor implements AccountRepository.
func (repo *AccountRepository) AccountRegistration(tx *sql.Tx, actor entity.Account) (int64, error) {
	SQL := `
	INSERT INTO actors(username, password, role_id, is_active, is_verified) 
	VALUES (?, ?, ?, ?, ?)`
	varArgs := []interface{}{
		actor.Username,
		actor.Password,
		actor.RoleId,
		actor.IsActive,
		actor.IsVerified,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Login implements AccountRepository.
func (repo *AccountRepository) SignIn(tx *sql.Tx, token entity.Token) (string, error) {
	SQL := `INSERT INTO authentications(token) VALUES (?)`
	varArgs := []interface{}{
		token.Token,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return "error repository", err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return "error repository", err
	}

	return fmt.Sprintf("rows affected: %d", i), nil
}
