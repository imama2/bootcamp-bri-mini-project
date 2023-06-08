package customer

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AccountUseCase UseCaseCustomerInterface
}

func NewHandler(userUC UseCaseCustomerInterface) *Handler {
	return &Handler{
		AccountUseCase: userUC,
	}
}

func (h Handler) Route(app *gin.Engine) {
	// account
	a := NewCostumerRequestHandler(h.AccountUseCase)
	a.RouteHandler(app)
}
