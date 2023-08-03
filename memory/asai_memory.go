package memory

import (
	"github.com/tmc/langchaingo/schema"
)

type AsaiMemory struct {}

func NewMemory(dsn string) *AsaiMemory {
	return &AsaiMemory{}
}

func (memory *AsaiMemory) GetSessionID() string {
	return ""
}

func (memory *AsaiMemory) SetSessionID(id string) {

}

func (memory *AsaiMemory) Buffer() schema.Memory {
	return nil
}

