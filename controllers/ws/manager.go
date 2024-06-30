package ws

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/GoLangWebSDK/crud/database"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	sync.RWMutex
	clients map[*Client]bool
}

func NewManager(db *database.Database) *Manager {
	mng := &Manager{
		clients: make(map[*Client]bool),
	}

	return mng
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
		websocketUpgrader.CheckOrigin = func(r *http.Request) bool {
			return r.Header.Get("Origin") == "http://localhost:5173"
		}
	}

	clinetConn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to initate socket:", err)
		return
	}

	ragURL := "ws://localhost:8080"
	ragConn, _, err := websocket.DefaultDialer.Dial(ragURL, nil)
	if err != nil {
		fmt.Println("Failed to connect to RAG server:", err)
		return
	}

	client := NewClient(clinetConn, ragConn, m)
	m.addClient(client)

	ctx := context.Background()
	go client.ProxyConnection(ctx)
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.clientConn.Close()
		client.ragConn.Close()
		delete(m.clients, client)
	}
}
