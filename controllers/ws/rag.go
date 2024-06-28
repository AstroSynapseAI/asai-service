package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
)

const DefaultRagURL = "ws://asai-rag:8081/ws/chat"

type RAGClient struct {
	Conn *websocket.Conn
}

func NewRAGClient(url string) *RAGClient {
	if url == "" {
		url = DefaultRagURL
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Failed to connect to RAG server:", err)
		return nil
	}

	client := &RAGClient{
		Conn: conn,
	}

	return client
}

func (rag *RAGClient) RelayMsg(payload []byte, response func(chunk []byte)) {

	err := rag.Conn.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		fmt.Println("Failed to send message to RAG server:", err)
		return
	}

	for {
		_, chunk, err := rag.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message from RAG server:", err)
			return
		}

		response(chunk)
	}
}
