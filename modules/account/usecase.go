package account

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
)

type UseCaseAccount interface {
	CreateAccount(user AccountParam) (entities.Account, error)
	GetAccountByID(id uint) (entities.Account, error)
	GetAccountByUsernameAndPassword(user AccountParam) (entities.Account, error)
}

type useCaseAccount struct {
	accountRepo repositories.AccountRepositoryInterface
}

func (uc useCaseAccount) CreateAccount(account AccountParam) (entities.Account, error) {
	var newAccount *entities.Account

	newAccount = &entities.Account{
		Username:   account.Username,
		Password:   account.Password,
		RoleId:     account.RoleId,
		Isverified: account.IsVerified,
		Isactive:   account.IsActive,
	}

	_, err := uc.accountRepo.CreateAccount(newAccount)
	if err != nil {
		return *newAccount, err
	}
	return *newAccount, nil
}

func (uc useCaseAccount) GetAccountByID(id uint) (entities.Account, error) {
	var cust entities.Account
	cust, err := uc.accountRepo.GetAccountByID(id)
	return cust, err
}

func (uc useCaseAccount) GetAccountByUsernameAndPassword(username, password string) (entities.Account, error) {
	var cust entities.Account
	cust, err := uc.accountRepo.GetAccountByUsernameAndPassword(username, password)
	return cust, err
}
