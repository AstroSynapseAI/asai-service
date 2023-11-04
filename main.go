package main

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/app"
	"github.com/AstroSynapseAI/app-service/app/server"
)

func main() {
	app := app.NewConfig()

	app.InitDB()

	asaiServer := server.NewAsaiServer()

	err := app.RunServer(asaiServer)
	if err != nil {
		fmt.Println("Failed to run server:", err)
		return
	}
}
