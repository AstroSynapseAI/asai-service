package agents

import (
	"github.com/tmc/langchaingo/schema"
)

type SearchAgentOptions func(agent *SearchAgent)

func WithMemory(memory schema.Memory) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Memory = memory
	}
}

func applySearchOptions(options ...SearchAgentOptions) *SearchAgent {
	searchAgent := &SearchAgent{}

	for _, option := range options {
		option(searchAgent)
	}

	return searchAgent
}