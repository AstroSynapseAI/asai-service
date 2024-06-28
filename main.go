package main

import (
	"fmt"

	"github.com/AstroSynapseAI/asai-service/app"
	"github.com/AstroSynapseAI/asai-service/app/server"
)

func main() {
	app := app.NewConfig()

	app.LoadEnvironment()

	app.InitDB()

	asaiServer := server.NewAsaiServer()

	err := app.RunServer(asaiServer)
	if err != nil {
		fmt.Println("Failed to run server:", err)
		return
	}
}
