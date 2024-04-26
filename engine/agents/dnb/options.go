package dnb

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms"
)

type DNBAgentOptions func(*DNBAgent)

type config struct {
	DNBAPISecret string `json:"dnb_api_secret,omitempty"`
	DNBAPIKey    string `json:"dnb_api_key,omitempty"`
	DNBAPIToken  string `json:"dnb_api_token,omitempty"`
}

func WithPrimer(primer string) DNBAgentOptions {
	return func(dnbAgent *DNBAgent) {
		dnbAgent.Primer = primer
	}
}

func WithLLM(llm llms.Model) DNBAgentOptions {
	return func(dnbAgent *DNBAgent) {
		dnbAgent.LLM = llm
	}
}

func WithConfig(data string) DNBAgentOptions {
	return func(dnbAgent *DNBAgent) {
		var configData config

		err := json.Unmarshal([]byte(data), &configData)
		if err != nil {
			fmt.Println("Error decoding config data:", err)
			return
		}

		dnbAgent.Config = configData
	}
}
