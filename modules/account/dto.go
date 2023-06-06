package account

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
	"github.com/imama2/bootcamp-bri-mini-project/entities"
)

type AccountParam struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleId     uint   `json:"roleId"`
	IsVerified bool   `json:"isVerified"`
	IsActive   bool   `json:"isActive"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data AccountParam `json:"data"`
}

type FindAccount struct {
	dto.ResponseMeta
	Data entities.Account `json:"data"`
}
