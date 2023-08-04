package memory

import (
	"github.com/AstroSynapseAI/engine-service/memory"
	lc_memory "github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
)

type AsaiMemory struct {
	buffer schema.Memory
}

func NewMemory(dsn string) *AsaiMemory {
	chatHistory := memory.NewPersistentChatHistory(dsn)
	buffer := lc_memory.NewConversationBuffer(lc_memory.WithChatHistory(chatHistory))

	return &AsaiMemory{
		buffer: buffer,
	}
}

func (memory *AsaiMemory) GetSessionID() string {
	return ""
}

func (memory *AsaiMemory) SetSessionID(id string) {

}

func (memory *AsaiMemory) Buffer() schema.Memory {
	return nil
}

