package user

import (
	"github.com/imama2/bootcamp-bri-mini-project/dto"
)

type Controller struct {
	uc UsecaseInterface
}

type ControllerInterface interface {
	GetUserByID(payload Payload) dto.Response
}
