package main

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
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

	router := user.NewRouter()
	router.Route(request)

}
