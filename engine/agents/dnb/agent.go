package dnb

import (
	"context"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

type DNBAgent struct {
	Primer   string
	LLM      llms.Model
	Executor agents.Executor
	Config   config
}

var _ tools.Tool = &DNBAgent{}

func NewDNBAgent(options ...DNBAgentOptions) (*DNBAgent, error) {
	dnbAgent := &DNBAgent{}

	for _, option := range options {
		option(dnbAgent)
	}

	return dnbAgent, nil
}

func (DNBAgent) Name() string {
	return "DNB"
}

func (DNBAgent) Description() string {
	return "DNB agent"
}

func (DNBAgent) Call(ctx context.Context, input string) (string, error) {
	return "", nil
}
