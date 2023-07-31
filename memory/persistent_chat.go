package memory

import "github.com/tmc/langchaingo/schema"

type PersistentChatHistory struct {
	
}

var _ schema.ChatMessageHistory = &PersistentChatHistory{}

func NewPersistentChatHistory() *PersistentChatHistory {
	return &PersistentChatHistory{}
}

func (history *PersistentChatHistory) Messages() ([]schema.ChatMessage, error) {
	return []schema.ChatMessage{}, nil
}

func (history *PersistentChatHistory) AddAIMessage(message string) error {
	return nil
}

func (history *PersistentChatHistory) AddUserMessage(message string) error {
	return nil
}

func (history *PersistentChatHistory) AddMessage(message schema.ChatMessage) error {
	return nil	
}

func (history *PersistentChatHistory) SetMessages(messages []schema.ChatMessage) error {
	return nil
}

func (history *PersistentChatHistory) Clear() error {
	return nil
}
