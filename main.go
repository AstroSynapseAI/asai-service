package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/AstroSynapseAI/engine-service/agents"
	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/AstroSynapseAI/engine-service/memory"
	"github.com/GoLangWebSDK/rest"
	lc_chains "github.com/tmc/langchaingo/chains"
)

var asaiMemory *memory.AsaiMemory
var asaiChain  *chains.AsaiChain

func init() {
	config.LoadEnvironment()

	dsn := config.SetupPostgreDSN()
	asaiMemory = memory.NewMemory(dsn)

	searchAgent, err := agents.NewSearchAgent(
		agents.WithMemory(asaiMemory.Buffer()),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	asaiChain = chains.NewAsaiChain(
		chains.WithMemory(asaiMemory),
		chains.WithSearchAgent(searchAgent),
	)
}

func main() {
	fmt.Println("main")
	
	router := rest.NewRouter()
	ctrl := rest.NewController(router)

	ctrl.Post("/api/chat/msg", func(ctx *rest.Context) {
		ctx.SetContentType("application/json")
		
		// Parse the incoming http request
		var request struct {
			SessionId  string `json:"session_id"`
			UserPrompt string `json:"user_prompt"`
			History    string `json:"history"`
		}

		err := ctx.JsonDecode(&request)
		if err != nil {
			fmt.Println("Bad Request: %v", err)
			ctx.JsonResponse(400, err)
			return
		}

		asaiMemory.SetSessionID(request.SessionId)

		response, err := lc_chains.Run(context.Background(), asaiChain, request.UserPrompt)
		if err != nil {
			fmt.Println(err)
			ctx.JsonResponse(500, err)
			return
		} 

		ctx.JsonResponse(200, response)
	})

	router.Mux.StrictSlash(true)

	port := os.Getenv("PORT")

	if port == "" {
		router.Listen(":8082")
		return
	}

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// Start the HTTP server using the router and the listener
	err = http.Serve(listener, router.Mux)

	if err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}