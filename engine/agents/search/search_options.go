package search

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
)

type SearchAgentOptions func(agent *SearchAgent)

func WithMemory(memory schema.Memory) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Memory = memory
	}
}

func WithPrimer(primer string) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Primer = primer
	}
}

func WithLLM(llm llms.LanguageModel) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.LLM = llm
	}
}

func WithTools(tools []tools.Tool) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Tools = tools
	}
}

