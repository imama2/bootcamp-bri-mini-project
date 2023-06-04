package main

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer"
)

func main() {
	var request = dto.Request{
		Body: map[string]string{
			"id": "1",
		},
		Method: "GET",
		Path:   "/get-customer",
		Header: map[string]string{
			"Authorization": "token",
		},
	}
	router := customer.NewRouter()
	router.Route(request)
}
