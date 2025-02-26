package main

import (
	"community/back/config"
	"community/back/router"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Port)

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
