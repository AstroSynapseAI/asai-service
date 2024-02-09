package chains

import (
	"context"
	"fmt"
	"time"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/AstroSynapseAI/app-service/engine/callbacks"
	"github.com/AstroSynapseAI/app-service/engine/memory"
	"github.com/AstroSynapseAI/app-service/engine/templates"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"

	asaiTools "github.com/AstroSynapseAI/app-service/engine/tools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
)

// Prompt for intializing conversation
const InitiativePrompt = "New user, has connected."

type AsaiChain struct {
	LLM        llms.LanguageModel
	Memory     *memory.AsaiMemory
	Agents     []tools.Tool
	Stream     func(context.Context, []byte)
	ClientType string
	config     engine.AvatarConfig
}

func NewAsaiChain(db *database.Database) *AsaiChain {
	asaiChain := &AsaiChain{
		config: engine.NewConfig(db),
		Memory: &memory.AsaiMemory{},
		Agents: []tools.Tool{},
	}

	return asaiChain

	// create search agent
	// searchAgent, err := search.NewSearchAgent()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	// // create browser agent
	// scraperAgent, err := browser.New()
	// if err != nil {
	// 	return nil, err
	// }

	// // create library agent
	// // currently using a simple tool for extracting documents
	// libraryAgent, err := documents.NewLoader()
	// if err != nil {
	// 	return nil, err
	// }

	// return &AsaiChain{
	// 	LLM:    config.GetAvatarLLM(),
	// 	Memory: asaiMemory,
	// 	config: config,
	// 	Agents: []tools.Tool{
	// 		searchAgent,
	// 		scraperAgent,
	// 		libraryAgent,
	// 	},
	// }, nil
}

func (chain *AsaiChain) LoadAvatar(userID uint, sessionID string, clientType string) {
	chain.config.LoadConfig(userID)
	chain.LLM = chain.config.GetAvatarLLM()
	chain.ClientType = clientType
	chain.Memory = memory.NewMemory(chain.config)
	chain.Memory.SetSessionID(sessionID)
}

func (chain *AsaiChain) SetStream(stream func(context.Context, []byte)) {
	chain.Stream = stream
}

func (chain *AsaiChain) SetSessionID(id string) {
	chain.Memory.SetSessionID(id)
}

func (chain *AsaiChain) SetClientType(clientType string) {
	chain.ClientType = clientType
}

func (chain *AsaiChain) LoadHistory() []schema.ChatMessage {
	return chain.Memory.Messages()
}

func (chain *AsaiChain) Prompt(ctx context.Context, input string) (string, error) {
	asaiAgent := agents.NewConversationalAgent(
		chain.LLM,
		chain.Agents,
	)

	obChain, err := NewOnboardingChain(chain.config, chain.Memory)
	if err != nil {
		return "", err
	}

	obResponse, err := obChain.Call(ctx, input)
	if err != nil {
		return "", err
	}

	tmplt := chain.loadTemplate(map[string]interface{}{
		"onboarding": obResponse,
	})

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

func (chain *AsaiChain) Run(ctx context.Context, input string, options ...chains.ChainCallOption) error {
	fmt.Println("Asai Chain Running...")

	// need to try this might be I initally loaded the prompt option wrong in the Executor
	// asaiAgent := agents.NewConversationalAgent(llm, chain.Agents, agents.WithPrompt(promptTmplt))

	agentCallback := callbacks.NewStreamHandler()
	agentCallback.ReadFromEgress(ctx, chain.Stream)

	asaiAgent := agents.NewConversationalAgent(
		chain.LLM,
		chain.Agents,
		agents.WithCallbacksHandler(agentCallback),
	)

	tmplt := chain.loadTemplate(map[string]interface{}{})

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

func (chain *AsaiChain) loadTemplate(values map[string]any) prompts.PromptTemplate {
	// load Asai persona prompt template
	template, err := templates.Load("default_primer.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Default Primer: ", template)

	// create agent prompt template
	return prompts.PromptTemplate{
		Template:       template,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]interface{}{
			"avatar_name":        chain.config.GetAvatarName(),
			"primer":             chain.config.GetAvatarPrimer(),
			"agent_names":        asaiTools.Names(chain.Agents),
			"agent_descriptions": asaiTools.Descriptions(chain.Agents),
			"date":               time.Now().Format("January 02, 2006"),
			"client_type":        chain.ClientType,
			"history":            "",
		},
	}
}
