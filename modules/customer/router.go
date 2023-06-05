package customer

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

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterCustomer {
	return RouterCustomer{CustomerRequestHandler: NewCostumerRequestHandler(
		dbCrud,
	)}
}

func (r RouterCustomer) Handle(router *gin.Engine) {
	basepath := "/customer"
	user := router.Group(basepath)

	user.POST("/register",
		r.CustomerRequestHandler.CreateCustomer,
	)

	user.GET("/register/:id",
		r.CustomerRequestHandler.GetCustomerByID)
}
