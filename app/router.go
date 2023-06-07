package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
)

func RouterInitiate(app *gin.Engine, DB *sql.DB) {
	// account setup
	accountRepo := repositories.NewAccountRepository()
	accountUC := account.NewAccountUseCase(accountRepo, DB)
	AccountHandler := account.NewHandler(accountUC)
	AccountHandler.Route(app)

	// Customer setup
	customerRepo := customerRepo.NewCustomerRepository()
	customerUC := customerUC.NewCustomerUseCase(customerRepo, DB)
	CustomerHandler := customerHandler.NewHandler(customerUC)
	CustomerHandler.Route(app)
}
