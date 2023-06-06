package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer"
	"github.com/imama2/bootcamp-bri-mini-project/utils/db"
)

func main() {
	router := gin.New()
	dbCrud := db.GormMysql()

	customerHandler := customer.NewRouter(dbCrud)
	customerHandler.Handle(router)

	accountHandler := account.NewRouter(dbCrud)
	accountHandler.Handle(router)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
