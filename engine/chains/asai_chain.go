package chains

import (
	"context"
	"fmt"
	"time"

	config "github.com/AstroSynapseAI/app"
	"github.com/AstroSynapseAI/engine/agents/browser"
	"github.com/AstroSynapseAI/engine/agents/search"
	"github.com/AstroSynapseAI/engine/memory"
	"github.com/AstroSynapseAI/engine/templates"
	"github.com/AstroSynapseAI/engine/tools/documents"

	asaiTools "github.com/AstroSynapseAI/engine/tools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
)

type AsaiChain struct {
	LLM    *openai.Chat
	Memory *memory.AsaiMemory
	Agents []tools.Tool
	Stream func(context.Context, []byte)
}

func NewAsaiChain() (*AsaiChain, error) {
	dsn := config.SetupPostgreDSN()
	asaiMemory := memory.NewMemory(dsn)

	// create llm
	llm, err := openai.NewChat(
		openai.WithModel("gpt-4"),
	)
	if err != nil {
		return nil, err
	}

	// create search agent
	searchAgent, err := search.NewSearchAgent()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create browser agent
	scraperAgent, err := browser.New()
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
		LLM:    llm,
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

func (chain AsaiChain) LoadHistory() []schema.ChatMessage {
	return chain.Memory.Messages()
}

func (chain AsaiChain) Prompt(ctx context.Context, input string) (string, error) {
	fmt.Println("Asai Prompt Running...")

	chain.loadTemplate()

	asaiAgent := agents.NewConversationalAgent(
		chain.LLM,
		chain.Agents,
	)

	tmplt := chain.loadTemplate()
	asaiAgent.Chain = chains.NewLLMChain(chain.LLM, tmplt)

	executor := agents.NewExecutor(
		asaiAgent,
		chain.Agents,
		agents.WithMemory(chain.Memory.Buffer()),
	)

	response, err := chains.Run(ctx, executor, input)
	if err != nil {
		return "", err
	}

	return response, nil

}

func (chain AsaiChain) Run(ctx context.Context, input string, options ...chains.ChainCallOption) error {
	fmt.Println("Asai Chain Running...")

	// need to try this might be I initally loaded the proompt option wrong in the Executor
	// asaiAgent := agents.NewConversationalAgent(llm, chain.Agents, agents.WithPrompt(promptTmplt))

	agentCallback := callbacks.NewFinalStreamHandler()
	agentCallback.ReadFromEgress(chain.Stream)

	asaiAgent := agents.NewConversationalAgent(
		chain.LLM,
		chain.Agents,
		agents.WithCallbacksHandler(agentCallback),
	)

	tmplt := chain.loadTemplate()

	asaiAgent.Chain = chains.NewLLMChain(chain.LLM, tmplt)

	executor := agents.NewExecutor(
		asaiAgent,
		chain.Agents,
		agents.WithMemory(chain.Memory.Buffer()),
		agents.WithCallbacksHandler(agentCallback),
	)

	// run the agent
	_, err := chains.Run(ctx, executor, input, options...)
	if err != nil {
		return err
	}

	return nil
}

func (chain AsaiChain) loadTemplate() prompts.PromptTemplate {
	// load Asai persona prompt template
	template, err := templates.Load("persona.txt")
	if err != nil {
		fmt.Println(err)
	}

	// create agent prompt template
	return prompts.PromptTemplate{
		Template:       template,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]interface{}{
			"agent_names":        asaiTools.Names(chain.Agents),
			"agent_descriptions": asaiTools.Descriptions(chain.Agents),
			"date":               time.Now().Format("January 02, 2006"),
			"history":            "",
		},
	}
}
