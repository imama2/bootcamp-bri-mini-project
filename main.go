package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imama2/bootcamp-bri-mini-project/app"
	"github.com/imama2/bootcamp-bri-mini-project/utils/config"
	"github.com/imama2/bootcamp-bri-mini-project/utils/db"
	"log"
)

func main() {
	configdb, err := config.LoadConfig()
	if err != nil {
		log.Printf("Config load error : ", err)
	}
	DB := db.NewDB(configdb)

	req := gin.Default()
	app.RouterInitiate(req, DB)

	req.Run(configdb.Server.Port)

}
