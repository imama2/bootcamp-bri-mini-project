package customer

import (
	"github.com/imama2/bootcamp-bri-mini-project/entities"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
)

type Usecase struct {
	userRepo repositories.UserRepositoryInterface
}

type UsecaseInterface interface {
	GetUserByID(payload Payload) []entities.Customer
}

func (uc Usecase) GetUserByID(payload Payload) []entities.Customer {
	user := uc.userRepo.GetByID(payload.id)

	// if len customer == 0 return no customer
	if len(user) == 0 {
		return nil
	}

	return user
}
