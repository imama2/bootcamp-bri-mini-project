package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
)

func RouterInitiate(app *gin.Engine, DB *sql.DB) {
	// account setup
	accountRepo := repositories.NewAccountRepository()
	accountUC := account.NewAccountUseCase(accountRepo, DB)
	AccountHandler := account.NewAccountRequestHandler(accountUC)
	AccountHandler.RouteHandler(app)

	customerRepo := repositories.NewCustomerRepository()
	customerUC := customer.NewCustomerUseCase(customerRepo, DB)
	CustomerHandler := customer.NewCostumerRequestHandler(customerUC)
	CustomerHandler.RouteHandler(app)
}
