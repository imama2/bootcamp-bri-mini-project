package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer"
	"github.com/imama2/bootcamp-bri-mini-project/utils/db"
)

func main() {
	router := gin.New()
	dbCrud := db.GormMysql()

	//checkdb, err := dbCrud.DB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////ping to database
	//errconn := checkdb.Ping()
	//if err != nil {
	//	log.Fatal(errconn)
	//}
	//
	//fmt.Println("database connected")
	//
	//errRouter := router.Run(":8080")
	//if errRouter != nil {
	//	fmt.Println("error running server", errRouter)
	//	return
	//}
	customerHandler := customer.NewRouter(dbCrud)
	customerHandler.Handle(router)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
