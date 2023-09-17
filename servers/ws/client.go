package ws

import (
	"encoding/json"
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
	// The reason why it has to be less than PingRequency is becuase otherwise it will send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	sessionID  string
	egress     chan []byte
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		egress:     make(chan []byte),
		connection: conn,
		manager:    manager,
	}
}

func (client *Client) ReadMsgs() {
	defer func() {
		client.manager.removeClient(client)
	}()

	// Set Max Size of Messages in Bytes
	// client.connection.SetReadLimit(512)

	// Configure Wait time for Pong response, use Current time + pongWait
	// This has to be done here to set the first initial timer.
	if err := client.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	// Configure how to handle Pong responses
	client.connection.SetPongHandler(client.pongHandler)

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

		var response string = dummyResponseString()

		for c := range client.manager.clients {
			if c.sessionID == client.sessionID {
				c.egress <- []byte(response)
			}
		}

	}
}

func (client *Client) SendMsgs() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		client.manager.removeClient(client)
	}()

	for {
		select {
		case msg, ok := <-client.egress:
			if !ok {
				if err := client.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					fmt.Println("Connection closed:", err)
				}
				return
			}

			// Write a Regular text message to the connection
			if err := client.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("Connection closed:", err)
			}
			fmt.Println("Msg sent:", msg)
		case <-ticker.C:
			fmt.Println("Ping...")

			if err := client.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				fmt.Println("Connection closed:", err)
				return
			}
		}
	}

}

func (client *Client) pongHandler(pongMsg string) error {
	// Current time + Pong Wait time
	fmt.Println("Pong...")
	return client.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func dummyResponseString() string {
	return "dummy response"
}
