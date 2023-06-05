package customer

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type controllerCustomer struct {
	customerUseCase UseCaseCustomer
}

type ControllerCustomer interface {
	CreateCustomer(req CustomerParam) (any, error)
	GetCustomerByID(id uint) (FindCustomer, error)
}

func (uc controllerCustomer) CreateCustomer(req CustomerParam) (any, error) {
	cust, err := uc.customerUseCase.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Firstname: cust.Firstname,
			Lastname:  cust.Lastname,
			Email:     cust.Email,
			Avatar:    cust.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) GetCustomerByID(id uint) (FindCustomer, error) {
	var res FindCustomer
	user, err := uc.customerUseCase.GetCustomerByID(id)
	if err != nil {
		return FindCustomer{}, err
	}
	res.Data = user
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success Update user",
		Message:      "Success Register",
		ResponseTime: "",
	}
	return res, nil

}
