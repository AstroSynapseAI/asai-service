package agents

import (
	"github.com/tmc/langchaingo/schema"
)

type SearchAgentOptions func(agent *SearchAgent)

func WithMemory(memory schema.Memory) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.memory = memory
	}
}