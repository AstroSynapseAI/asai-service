package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/GoLangWebSDK/rest"
)

var asaiChain  *chains.AsaiChain

func init() {
	var err error
	config.LoadEnvironment()

	asaiChain, err = chains.NewAsaiChain()
	if err != nil {
		fmt.Println("Failed to create AsaiChain:", err)
		return
	}
}

func main() {
	router := rest.NewRouter()
	ctrl := rest.NewController(router)

	ctrl.Post("/api/chat/msg", func(ctx *rest.Context) {
		ctx.SetContentType("application/json")
		
		// Parse the incoming http request
		var request struct {
			SessionId  string `json:"session_id"`
			UserPrompt string `json:"user_prompt"`
		}

		err := ctx.JsonDecode(&request)
		if err != nil {
			fmt.Println("Bad Request: %w", err)
			ctx.JsonResponse(400, err)
			return
		}

		asaiChain.SetSessionID(request.SessionId)

		response, err := asaiChain.Run(context.Background(), request.UserPrompt)
		if err != nil {
			fmt.Println(err)
			ctx.JsonResponse(500, err)
			return
		}

		var responseJson struct {
			Content string `json:"content"`
		}

		responseJson.Content = response

		ctx.JsonResponse(200, responseJson)
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