package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/GoLangWebSDK/rest"
	"github.com/gorilla/websocket"
)

var asaiChain *chains.AsaiChain

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

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	router.Mux.StrictSlash(true)

	router.Mux.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })

		// Launch each client's message loop in its own goroutine.
		go func() {
			for {
				messageType, message, err := conn.ReadMessage()
				if err != nil {
					fmt.Println(err)
					return
				}

				var request struct {
					Action string `json:"action"`
					Data   struct {
						SessionId  string `json:"session_id"`
						UserPrompt string `json:"user_prompt"`
					} `json:"data"`
				}

				err = json.Unmarshal(message, &request)
				if err != nil {
					fmt.Println(err)
					return
				}

				if request.Action != "chat_message" {
					return
				}

				asaiChain.SetSessionID(request.Data.SessionId)
				asaiResponse, err := asaiChain.Run(context.Background(), request.Data.UserPrompt)
				if err != nil {
					fmt.Println(err)
					return
				}

				var response struct {
					Content string `json:"content"`
				}

				response.Content = asaiResponse

				responseJson, err := json.Marshal(&response)
				if err != nil {
					fmt.Println(err)
					return
				}

				// Hardcode the message type to websocket.TextMessage
				if err := conn.WriteMessage(messageType, responseJson); err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	})

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
