package chains

import (
	"context"
	"fmt"
	"time"

	"github.com/AstroSynapseAI/engine-service/agents/search"
	"github.com/AstroSynapseAI/engine-service/config"
	"github.com/AstroSynapseAI/engine-service/memory"
	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/AstroSynapseAI/engine-service/tools/documents"
	"github.com/AstroSynapseAI/engine-service/tools/scraper"

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

	// create browser agent
	// currently using the scrapper tool as an agent until the browsing
	// get's more complex. ie. introduction of headless browser
	scraperAgent, err := scraper.NewScraper()
	if err != nil {
		return nil, err
	}

	// create library agent
	// currently using a simple tool for extracting documents
	libraryAgent, err := documents.NewLoader()
	if err != nil {
		return nil, err
	}

	return &AsaiChain{
		Memory: asaiMemory,
		Agents: []tools.Tool{
			searchAgent,
			scraperAgent,
			libraryAgent,
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

	// create agent prompt template
	promptTmplt := prompts.PromptTemplate{
		Template:       template,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]interface{}{
			"agent_names":        asaiTools.Names(chain.Agents),
			"agent_descriptions": asaiTools.Descriptions(chain.Agents),
			"today":              time.Now().Format("January 02, 2006"),
			"history":            "",
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
	
	// run the agent
	response, err := chains.Run(ctx, executor, input)
	if err != nil {
		return "", err
	}

	return response, nil
}