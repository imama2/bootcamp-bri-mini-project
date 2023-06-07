package customer

import (
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/customer"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer/do"
)

func DTOCustomer(et entity.Customer) do.Customer {
	return do.Customer{
		ID:        et.ID,
		FirstName: et.FirstName,
		LastName:  et.LastName,
		Email:     et.Email,
		Avatar:    et.Avatar,
		CreatedAt: et.CreatedAt,
		UpdatedAt: et.UpdatedAt,
	}
}

func DTOListCustomer(et []entity.Customer) []do.Customer {
	result := make([]do.Customer, 0)
	for _, v := range et {
		result = append(result, DTOCustomer(v))
	}

	return result
}
