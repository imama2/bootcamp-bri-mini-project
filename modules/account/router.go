package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterAccount struct {
	AccountRequestHandler RequestHandlerAccount
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterAccount {
	return RouterAccount{AccountRequestHandler: NewAccountRequestHandler(
		dbCrud,
	)}
}

func (r RouterAccount) Handle(router *gin.Engine) {
	basepath := "/account"
	user := router.Group(basepath)

	user.POST("/register",
		r.AccountRequestHandler.CreateAccount,
	)

	user.GET("/login",
		r.AccountRequestHandler.GetAccountByUsernameAndPassword)
}
