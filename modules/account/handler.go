package account

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AccountUseCase UseCaseAccountInterface
}

func NewHandler(userUC UseCaseAccountInterface) *Handler {
	return &Handler{
		AccountUseCase: userUC,
	}
}

func (h Handler) Route(app *gin.Engine) {
	// account
	//b := NewAccountRequestHandler()
	_ = WebResponse{}
	_ = RequestHandlerAccount{AccountUseCase: h.AccountUseCase}
	a := NewAccountRequestHandler(h.AccountUseCase)
	a.RouteHandler(app)
}
