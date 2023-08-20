package chains

import (
	"context"
	"fmt"
	"time"

	"github.com/AstroSynapseAI/engine-service/agents/search"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/AstroSynapseAI/engine-service/memory"
	"github.com/AstroSynapseAI/engine-service/templates"

	asaiTools "github.com/AstroSynapseAI/engine-service/tools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

type AsaiChain struct {
	Memory 	 	*memory.AsaiMemory
	Agents 		[]tools.Tool	
}

func NewAsaiChain() (*AsaiChain, error) {
	dsn := config.SetupPostgreDSN()
	asaiMemory := memory.NewMemory(dsn)
	
	// create search agent
	searchAgent, err := search.NewSearchAgent()
	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create borowser agent

	// create librarian agent

	return &AsaiChain{
		Memory: asaiMemory,
		Agents: []tools.Tool{
			searchAgent,
		},
	}, nil
}

func (chain AsaiChain) SetSessionID(id string) {
	chain.Memory.SetSessionID(id)	
}

func (chain AsaiChain) Run(ctx context.Context, input string) (string, error) {
	// create llm handle
	llm, err := openai.NewChat(
		openai.WithModel("gpt-4"),
	)

	// load Asai persona prompt template
	template, err := templates.Load("persona.txt")
	if err != nil {
		return "", err
	}

	promptTmplt := prompts.PromptTemplate{
		Template:       template,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]interface{}{
			"tool_names":        asaiTools.Names(chain.Agents),
			"tool_descriptions": asaiTools.Descriptions(chain.Agents),
			"today":             time.Now().Format("January 02, 2006"),
			"history":           "",
		},
	}

	// create asai agent
	asaiAgent := agents.NewConversationalAgent(llm, chain.Agents)
	executor := agents.NewExecutor(
		asaiAgent, 
		chain.Agents,
		agents.WithPrompt(promptTmplt),	
		agents.WithMemory(chain.Memory.Buffer()),
	)
	
	response, err := chains.Run(ctx, executor, input)
	if err != nil {
		return "", err
	}

	return response, nil
}