package ws

import (
	"context"
	"fmt"
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
		ragClient:  NewRAGClient("ws://asai-rag:8081/ws/chat"),
	}
}

func (client *Client) MaintainConnection(ctx context.Context) {
	client.ragClient.Conn.SetPingHandler(func(msg string) error {
		return client.connection.WriteMessage(websocket.PingMessage, []byte(msg))
	})

	client.connection.SetPongHandler(func(msg string) error {
		return client.ragClient.Conn.WriteMessage(websocket.PongMessage, []byte(msg))
	})
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
