package ws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	// pongWait is how long we will await a pong response from client
	pongWait = 10 * time.Second
	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	// Because that can make decimals, so instead *9 / 10 to get 90%
	// The reason why it has to be less than PingRequency is becuase otherwise it will
	// send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	egress     chan []byte
	connection *websocket.Conn
	manager    *Manager
	ragClient  *RAGClient
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		egress:     make(chan []byte),
		connection: conn,
		manager:    manager,
		ragClient:  NewRAGClient("ws://localhost:8080/ws/chat"),
	}
}

func (client *Client) MaintainConnection(ctx context.Context) {
	ticker := time.NewTicker(pingInterval)

	// Configure Wait time for Pong response, use Current time + pongWait
	// This has to be done here to set the first initial timer.
	if err := client.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println("Failed to set read deadline:", err)
		return
	}

	client.connection.SetPongHandler(client.PongHandler)

	for {
		err := client.connection.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			fmt.Println("Ping failed:", err)
			return
		}
		// Wait for next tick
		<-ticker.C
	}
}

func (client *Client) PongHandler(pongMsg string) error {
	// Current time + Pong Wait time

	err := client.connection.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("Failed to set read deadline in pong handler:", err)
		return err
	}

	return nil
}

func (client *Client) ReadMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	for {
		_, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error reading message: %v", err)
			}
			break
		}

		go func() {
			client.ragClient.RelayMsg(payload, func(chunk []byte) {
				// Send message to egress
				client.egress <- chunk
			})
		}()
	}
}

func (client *Client) SendMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	for msg := range client.egress {
		// Write a Regular text message to the connection
		if err := client.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Sending message failed:", err)
			return
		}
	}
}
