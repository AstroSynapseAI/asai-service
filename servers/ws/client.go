package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AstroSynapseAI/engine-service/cortex/chains"
	"github.com/gorilla/websocket"
)

var (
	// pongWait is how long we will await a pong response from client
	pongWait = 30 * time.Second
	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	// Because that can make decimals, so instead *9 / 10 to get 90%
	// The reason why it has to be less than PingRequency is becuase otherwise it will
	// send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	sessionID  string
	egress     chan []byte
	connection *websocket.Conn
	manager    *Manager
	asaiChain  *chains.AsaiChain
}

func NewClient(conn *websocket.Conn, manager *Manager, chain *chains.AsaiChain) *Client {
	return &Client{
		egress:     make(chan []byte),
		connection: conn,
		manager:    manager,
		asaiChain:  chain,
	}
}

func (client *Client) MaintainConnection(ctx context.Context) {
	ticker := time.NewTicker(pingInterval)

	// Configure Wait time for Pong response, use Current time + pongWait
	// This has to be done here to set the first initial timer.
	if err := client.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	client.connection.SetPongHandler(client.pongHandler)

	for {
		fmt.Println("Ping...")
		err := client.connection.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
		// Wait for next tick
		<-ticker.C
	}
}

func (client *Client) ReadMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	// Set Max Size of Messages in Bytes
	// client.connection.SetReadLimit(512)

	for {
		_, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error reading message: %v", err)
			}
			break
		}

		var request struct {
			SessionId  string `json:"session_id"`
			UserPrompt string `json:"user_prompt"`
		}

		if err = json.Unmarshal(payload, &request); err != nil {
			fmt.Println("error marshalling message: ", err)
			break
		}

		client.sessionID = request.SessionId
		client.asaiChain.SetSessionID(request.SessionId)

		response, err := client.asaiChain.Run(context.Background(), request.UserPrompt)
		if err != nil {
			fmt.Println("error running chain: ", err)
			break
		}

		client.egress <- []byte(response)

		// Figuring out and testing LLM stream response.
		// The Langchain-go doesn't support streamed agent response atm
		// Need to make a contribution to Langchain-go
		// _, _ = client.asaiChain.Run(
		// 	ctx,
		// 	request.UserPrompt,
		// 	options.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		// 		fmt.Print(string(chunk))
		// 		client.egress <- chunk
		// 		return nil
		// 	}),
		// )
		// client.egress <- []byte("/end/")

		// var response string = dummyResponseString()
		// for c := range client.manager.clients {
		// 	if c.sessionID == client.sessionID {
		// 		for _, symbol := range getSymbols(response) {
		// 			c.egress <- []byte(symbol)
		// 		}
		// 		c.egress <- []byte("/end/")
		// 	}
		// }

	}
}

func (client *Client) SendMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	for msg := range client.egress {
		// Write a Regular text message to the connection
		if err := client.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Connection closed with error:", err)
			return
		}

		fmt.Println("Msg sent...")
	}
}

func (client *Client) pongHandler(pongMsg string) error {
	// Current time + Pong Wait time
	fmt.Println("Pong...")
	return client.connection.SetReadDeadline(time.Now().Add(pongWait))
}

// func dummyResponseString() string {
// 	return `Yes, I see the bug. The issue arises from the prefixPrinted variable inside the goroutine that's used to determine whether to print the "ASAI >" prefix. Once it's set to true for the first response, it never gets reset, so the prefix is not printed for subsequent responses.`
// }

// func getSymbols(s string) []string {
// 	var symbols []string
// 	for _, char := range s {
// 		symbols = append(symbols, string(char))
// 	}
// 	return symbols
// }
