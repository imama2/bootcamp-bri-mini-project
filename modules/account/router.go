package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//type Router struct {
//	rq RequestHandlerInterface
//}
//
//func NewRouter() Router {
//	return Router{}
//}
//
//func (r Router) Route(request dto.Request) {
//
//}

type RouterAccount struct {
	AccountRequestHandler RequestHandlerAccount
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterAccount {
	return RouterAccount{AccountRequestHandler: NewCostumerRequestHandler(
		dbCrud,
	)}
}

func (r RouterAccount) Handle(router *gin.Engine) {
	basepath := "/account"
	user := router.Group(basepath)

	user.POST("/register",
		r.AccountRequestHandler.CreateAccount,
	)

	user.GET("/register/:id",
		r.AccountRequestHandler.GetAccountByID)
}
