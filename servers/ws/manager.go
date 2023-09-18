package ws

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/AstroSynapseAI/engine-service/chains"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	// CheckOrigin:     checkOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	ctx context.Context
	sync.RWMutex
	clients map[*Client]bool
}

func NewManager(ctx context.Context) *Manager {
	mng := &Manager{
		clients: make(map[*Client]bool),
		ctx:     ctx,
	}

	return mng
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to initate socket:", err)
		return
	}

	asaiChain, _ := chains.NewAsaiChain()

	client := NewClient(conn, m, asaiChain)

	m.addClient(client)

	go client.ReadMsgs(m.ctx)
	go client.SendMsgs(m.ctx)

}

func (m *Manager) addClient(client *Client) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		// close connection
		client.connection.Close()
		// remove
		delete(m.clients, client)
	}
}
