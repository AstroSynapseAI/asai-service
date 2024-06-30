package ws

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	manager    *Manager
	clientConn *websocket.Conn
	ragConn    *websocket.Conn
}

func NewClient(clientConn *websocket.Conn, ragConn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		manager:    manager,
		clientConn: clientConn,
		ragConn:    ragConn,
	}
}

func (client *Client) ProxyConnection(ctx context.Context) {
	fmt.Println("Proxying connection on asai service!")

	defer client.manager.removeClient(client)

	go func() {
		for {
			messageTpe, payload, err := client.clientConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("error reading message: %v", err)
				}
				break
			}

			err = client.ragConn.WriteMessage(messageTpe, payload)
			if err != nil {
				fmt.Println("Failed to send message to RAG server:", err)
				return
			}
		}
	}()

	for {
		messageType, msg, err := client.ragConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error reading message: %v", err)
			}
			break
		}

		err = client.clientConn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Failed to send message to client:", err)
			return
		}
	}
}
