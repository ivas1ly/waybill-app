package main

import (
	"log"
	"os"

	"github.com/ivas1ly/waybill-app/config"
	"github.com/ivas1ly/waybill-app/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()
	app.Run(port)
}
