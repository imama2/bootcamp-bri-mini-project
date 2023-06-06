package repositories

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccount(dbCrud *gorm.DB) Account {
	return Account{
		db: dbCrud,
	}

}

type AccountRepositoryInterface interface {
	//GetAccountByID(id uint) (entities.Account, error)
	UpdateAccount(user *entities.Account) (any, error)
	CreateAccount(account *entities.Account) (*entities.Account, error)
	DeleteAccount(email string) (any, error)
	GetAccountByUsernameAndPassword(username, password string) (any, error)
}

//func (repo Account) GetAccountByID(id uint) (entities.Account, error) {
//	var account entities.Account
//	repo.db.First(&account, `id = ?`, id)
//	return account, nil
//}

func (repo Account) CreateAccount(account *entities.Account) (*entities.Account, error) {
	err := repo.db.Model(&entities.Account{}).Create(account).Error
	return account, err
}

func (repo Account) UpdateAccount(user *entities.Account) (any, error) {
	err := repo.db.Model(&entities.Account{}).
		Save(user).Error
	return nil, err
}

// DeleteAccount by Id and email
func (repo Account) DeleteAccount(email string) (any, error) {
	err := repo.db.Model(&entities.Account{}).
		Where("email = ?", email).
		Delete(&entities.Account{}).
		Error
	return nil, err
}

func (repo Account) GetAccountByUsernameAndPassword(username, password string) (any, error) {
	err := repo.db.Model(&entities.Account{}).
		Where("username = ? and password = ?", username, password).
		Error
	return nil, err
}
