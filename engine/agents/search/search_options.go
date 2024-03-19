package search

import (
	"encoding/json"
	"fmt"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

type config struct {
	DDGIsActive    bool
	GoogleAPIToken string
	GoogleIsActive bool
	ExaAPIToken    string
	ExaIsActive    bool
}

type SearchAgentOptions func(agent *SearchAgent)

func WithConfig(data string) SearchAgentOptions {
	return func(agent *SearchAgent) {
		var configData config

		err := json.Unmarshal([]byte(data), &configData)
		if err != nil {
			fmt.Println("Error decoding search onfig data:", err)
			return
		}
		agent.Config = configData
	}
}

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

func WithToolsConfig(tools []engine.AgentToolConfig) SearchAgentOptions {
	return func(agent *SearchAgent) {
	}
}
