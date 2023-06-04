package main

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
	"github.com/imama2/bootcamp-bri-mini-project/modules/user"
)

func main() {
	var request = dto.Request{
		Body: map[string]string{
			"id": "1",
		},
		Method: "GET",
		Path:   "/get-user",
		Header: map[string]string{
			"Authorization": "token",
		},
	}

	router := user.NewRouter()
	router.Route(request)

}
