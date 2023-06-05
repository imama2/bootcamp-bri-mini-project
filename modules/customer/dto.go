package customer

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
	"github.com/imama2/bootcamp-bri-mini-project/entities"
)

type CustomerParam struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entities.Customer `json:"data"`
}
