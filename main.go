package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/utils/db"
	"log"
)

func main() {
	//var request = dto.Request{
	//	Body: map[string]string{
	//		"id": "1",
	//	},
	//	Method: "GET",
	//	Path:   "/get-customer",
	//	Header: map[string]string{
	//		"Authorization": "token",
	//	},
	//}
	//router := customer.NewRouter()
	//router.Route(request)

	router := gin.New()
	dbCrud := db.GormMysql()

	checkdb, err := dbCrud.DB()
	if err != nil {
		log.Fatal(err)
	}

	//ping to database
	errconn := checkdb.Ping()
	if err != nil {
		log.Fatal(errconn)
	}

	fmt.Println("database connected")

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
