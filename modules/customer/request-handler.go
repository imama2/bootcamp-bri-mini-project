package customer

import (
	"fmt"

	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	GetUserByID(request dto.Request) dto.Response
}

func (rq RequestHandler) GetUserByID(request dto.Request) dto.Response {

	// convert response ke payload, terjadi validasi
	payload := Payload{
		ID: 1,
	}

	response := rq.ctrl.GetUserByID(payload)

	fmt.Println(response)
	return response
}
