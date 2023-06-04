package user

import "github.com/imama2/bootcamp-bri-mini-project/dto"

type Router struct {
	rq RequestHandlerInterface
}

func NewRouter() Router {
	return Router{}
}

func (r Router) Route(request dto.Request) {

}
