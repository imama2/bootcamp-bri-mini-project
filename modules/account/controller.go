package account

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type controllerAccount struct {
	accountUseCase UseCaseAccount
}

type ControllerAccount interface {
	CreateAccount(req AccountParam) (any, error)
	//GetAccountByID(id uint) (FindAccount, error)
	GetAccountByUsernameAndPassword(username, password string) (FindAccount, error)
}

func (uc controllerAccount) CreateAccount(req AccountParam) (any, error) {
	acc, err := uc.accountUseCase.CreateAccount(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create Account",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AccountParam{
			Username:   acc.Username,
			Password:   acc.Password,
			RoleId:     acc.RoleId,
			IsVerified: acc.Isverified,
			IsActive:   acc.Isactive,
		},
	}
	return res, nil
}

//func (uc controllerAccount) GetAccountByID(id uint) (FindAccount, error) {
//	var res FindAccount
//	user, err := uc.accountUseCase.GetAccountByID(id)
//	if err != nil {
//		return FindAccount{}, err
//	}
//	res.Data = user
//	res.ResponseMeta = dto.ResponseMeta{
//		Success:      true,
//		MessageTitle: "Success Update user",
//		Message:      "Success Register",
//		ResponseTime: "",
//	}
//	return res, nil
//
//}

func (uc controllerAccount) GetAccountByUsernameAndPassword(username, password string) (FindAccount, error) {
	var res FindAccount
	user, err := uc.accountUseCase.GetAccountByUsernameAndPassword(username, password)
	if err != nil {
		return FindAccount{}, err
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
