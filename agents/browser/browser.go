package browser

import (
	"context"

	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/AstroSynapseAI/engine-service/tools/scraper"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"

	asaiTools "github.com/AstroSynapseAI/engine-service/tools"
)

var _ tools.Tool = &BrowserAgent{} 

type BrowserAgent struct {
	Memory schema.Memory
	Executor agents.Executor
} 

func New(options ...BrowserAgentOptions) (*BrowserAgent, error) {
	browserAgent := &BrowserAgent{
		Memory: memory.NewSimple(),
	}

	for _, option := range options {
		option(browserAgent)
	}

	scraper, err := scraper.New()
	if err != nil {
		return nil, err
	}

	browserTools := []tools.Tool{scraper}

	browserTmplt, err := templates.Load("name.txt")	
	if err != nil {
		return nil, err
	}

	promptTmplt := prompts.PromptTemplate{
		Template:       browserTmplt,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"tool_names":        asaiTools.Names(browserTools),
			"tool_descriptions": asaiTools.Descriptions(browserTools),
			"history":           "",
		},
	}

	llm, err := openai.New(openai.WithModel("gpt-4"))
	if err != nil {
		return nil, err
	}

	agent := agents.NewOneShotAgent(
		llm, 
		browserTools, 
		agents.WithMemory(browserAgent.Memory),
		agents.WithPrompt(promptTmplt),
	)
	
	browserAgent.Executor = agents.NewExecutor(agent, browserTools)
	
	return browserAgent, nil
}

func (agent *BrowserAgent) Call(ctx context.Context, input string) (string, error) {
	return chains.Run(ctx, agent.Executor, input)
} 

func (agent *BrowserAgent) Name() string {
	return "Web Browser Agent"	
}

func (agent *BrowserAgent) Description() string {
	return `
		Web Browser Agent is an agent specialized in scraping and reading the web pages.
	`
}