package account

import (
	"database/sql"
	"errors"
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/account"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"github.com/imama2/bootcamp-bri-mini-project/package/helper"
	"github.com/imama2/bootcamp-bri-mini-project/package/security"
	"github.com/imama2/bootcamp-bri-mini-project/package/token"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"sync"
	"time"
)

type UseCaseAccountInterface interface {
	// Autentikasi
	AccountAuthentication(req do.Account) (do.ResToken, error) // generate token jwt
	// Registrasi akun
	AccountRegistration(req do.Account) (int64, error)
	// Admin Fetch
	GetAllAdmin(req do.Account, pagi do.Pagination) (do.ListActorWithPaging, error)

	// Super admin
	GetAllApprovalAdmin() ([]do.Approval, error)
	UpdateAdminStatusByID(reqReg do.Approval, reqActor do.Account) (int64, error)
	DeleteAdminByID(req do.Account) (int64, error)
}

func NewAccountUseCase(AccountRepository repositories.AccountRepositoryInterface, DB *sql.DB) UseCaseAccountInterface {
	return &UseCaseAccount{
		AccountRepository: AccountRepository,
		DB:                DB,
	}
}

type UseCaseAccount struct {
	AccountRepository repositories.AccountRepositoryInterface
	DB                *sql.DB
}

func (uc *UseCaseAccount) GetAllAdmin(req do.Account, pagi do.Pagination) (do.ListActorWithPaging, error) {
	var (
		err       error
		wg        sync.WaitGroup
		resPaging entity.Pagination
		result    []entity.Account
	)

	chListAdmin := make(chan []entity.Account, 1)
	chPaging := make(chan entity.Pagination, 1)
	errListAdmin := make(chan error, 1)
	errPagination := make(chan error, 1)

	// ?: Error tx with go routine, temporary solution using db queries, maybe tx MySQL doesn't support query select rows on goroutines
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
	et := entity.Account{
		Username: req.Username,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := uc.AccountRepository.GetAllAdmin(uc.DB, et, etPaging)
		if err != nil {
			errListAdmin <- err
		}
		chListAdmin <- result
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Get Total Data
		resPaging, err := uc.AccountRepository.Pagination(uc.DB, etPaging)
		if err != nil {
			errPagination <- err
		}
		chPaging <- resPaging
	}()
	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case result = <-chListAdmin:
			continue
		case resPaging = <-chPaging:
			continue
		case err = <-errListAdmin:
			return do.ListActorWithPaging{}, err
		case err = <-errPagination:
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
		Admins:     DTOAccountList(result),
	}

	return combineRes, nil
}

// DeleteAdminByID implements AccountUseCase.
func (uc *UseCaseAccount) DeleteAdminByID(req do.Account) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	entity := entity.Account{
		ID: req.ID,
	}
	result, err := uc.AccountRepository.DeleteAdminByID(tx, entity)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// UpdateAdminStatusByID implements AccountUseCase.
func (uc *UseCaseAccount) UpdateAdminStatusByID(reqReg do.Approval, reqActor do.Account) (int64, error) {
	var (
		wg     sync.WaitGroup
		result int64
	)
	chErr1 := make(chan error, 1)
	chErr2 := make(chan error, 1)
	chInt := make(chan int64, 1)
	chInt2 := make(chan int64, 1)

	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	etAdminReg := entity.Approval{
		AdminId: reqReg.AdminId,
		Status:  reqReg.Status,
	}
	etActor := entity.Account{
		ID:         reqActor.ID,
		IsVerified: reqActor.IsVerified,
		IsActive:   reqActor.IsActive,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// update admin_reg status only
		i, err := uc.AccountRepository.UpdateAdminRegStatusByAdminID(tx, etAdminReg)
		if err != nil {
			chErr1 <- err
		}
		chInt <- i
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// update actor is_verified & is_active
		i2, err := uc.AccountRepository.UpdateAdminStatusByAdminID(tx, etActor)
		if err != nil {
			chErr2 <- err
		}
		chInt2 <- i2
	}()
	wg.Wait()

	// get 2 channel data
	var totalRowsAffected int64
	for i := 0; i < 2; i++ {
		select {
		case result = <-chInt2:
			totalRowsAffected += result
		case result = <-chInt:
			totalRowsAffected += result
		case err = <-chErr1:
			return 0, err
		case err = <-chErr2:
			return 0, err
		}
	}
	return totalRowsAffected, nil
}

// GetAllApprovalAdmin implements AccountUseCase.
func (uc *UseCaseAccount) GetAllApprovalAdmin() ([]do.Approval, error) {
	result := make([]do.Approval, 0)
	tx, err := uc.DB.Begin()
	if err != nil {
		return result, err
	}
	defer helper.CommitOrRollback(err, tx)

	res, err := uc.AccountRepository.GetAllApprovalAdmin(tx)
	if err != nil {
		return result, err
	}

	return DTOListApprovalAdd(res), nil
}

// VerifyActorCredential implements AccountUseCase.
func (uc *UseCaseAccount) AccountAuthentication(req do.Account) (do.ResToken, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return do.ResToken{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	entity := entity.Account{
		Username: req.Username,
		Password: req.Password,
	}
	result, err := uc.AccountRepository.AccountAuthentication(tx, entity)
	if err != nil {
		return do.ResToken{}, err
	}

	// compare password
	isValid := security.CheckPasswordHash(req.Password, result.Password)
	if !isValid {
		return do.ResToken{}, errors.New("invalid username or password")
	}
	userDetail := DTOAccount(result)

	// generate token jwt
	// Create the Claims
	myClaims := token.AccountClaims{
		IDNum:      userDetail.ID,
		RoleID:     userDetail.RoleID,
		IsVerified: userDetail.IsVerified,
		IsActive:   userDetail.IsActive,
		ExpiresAt:  time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err := token.GenerateAccessToken(myClaims)
	if err != nil {
		return do.ResToken{}, err
	}

	res := do.ResToken{
		AccessToken: token,
	}
	return res, nil
}

// AddActor implements AccountUseCase.
func (uc *UseCaseAccount) AccountRegistration(req do.Account) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	// Hash Passwword
	hashPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	data := entity.Account{
		Username:   req.Username,
		Password:   hashPassword,
		RoleId:     1,
		IsActive:   false,
		IsVerified: false,
	}

	resultID, err := uc.AccountRepository.AccountRegistration(tx, data)
	if err != nil {
		return 0, err
	}

	adminReg := entity.Approval{
		AdminId:      resultID,
		SuperAdminId: 1,
		Status:       "pending",
	}

	result, err := uc.AccountRepository.ApprovalAdd(tx, adminReg)
	if err != nil {
		return 0, err
	}

	return result, nil
}
