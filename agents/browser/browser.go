package browser

import (
	"context"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
)

var _ tools.Tool = &BrowserAgent{} 

type BrowserAgent struct {
	memory schema.Memory
	executor agents.Executor
} 

func New(options ...BrowserAgentOptions) (*BrowserAgent, error) {
	browserAgent := &BrowserAgent{
		memory: memory.NewSimple(),
	}

	for _, option := range options {
		option(browserAgent)
	}

	return browserAgent, nil
}

func (agent *BrowserAgent) Call(ctx context.Context, input string) (string, error) {
	return chains.Run(ctx, agent.executor, input)
} 

func (agent *BrowserAgent) Name() string {
	return "Web Browser Agent"	
}

func (agent *BrowserAgent) Description() string {
	return `
		Web Browser Agent is an agent specialized in scraping and reading the web pages.
	`
}