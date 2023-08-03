package chains

import (
	"context"
	"time"

	"github.com/AstroSynapseAI/engine-service/agents"
	"github.com/AstroSynapseAI/engine-service/memory"
	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

type AsaiChain struct {
	Memory 	 	*memory.AsaiMemory
	SearchAgent	*agents.SearchAgent
	OutputKey	string
	InputKey	string
	
}

var _ chains.Chain = &AsaiChain{}

func NewAsaiChain(options ...ChainOptions) *AsaiChain {
	asaiChain := &AsaiChain{
		OutputKey: "output",
		InputKey: "input",
	}
	
	for _, option := range options {
		option(asaiChain)
	}

	return asaiChain
}

func (chain AsaiChain) Call(ctx context.Context, inputValues map[string]any, _ ...chains.ChainCallOption) (map[string]any, error) {
	// perform search
	userInput := inputValues["input"]
	searchResults, err := chains.Run(ctx, chain.SearchAgent.Executor(), userInput)

	// check if search was performed
	llm, err := openai.New(
		openai.WithModel("gpt-4"),
	)

	if err != nil {
		return nil, err
	}

	buffer, err := chain.Memory.Buffer().LoadMemoryVariables(inputValues)
	if err != nil {
		return nil, err
	}

	monitorTemplate, err := templates.Load("monitor.txt")
	if err != nil {
		return nil, err
	}

	monitorPrompt := prompts.NewPromptTemplate(monitorTemplate, []string{"input", "history", "searchResults"})
	monitor := chains.NewLLMChain(llm, monitorPrompt)
	monitorInput := map[string]interface{}{
		"input":  userInput,
		"history": buffer["history"],
		"searchResults": searchResults,
	}

	searchPerformed, err := chains.Predict(ctx, monitor, monitorInput)
	if err != nil {
		return nil, err
	}

	// answer the prompt
	personaTemplate, err := templates.Load("persona.txt")
	if err != nil {
		return nil, err
	}

	personaPrompt := prompts.NewPromptTemplate(personaTemplate, []string{
		"input", 
		"history", 
		"searchPerformed", 
		"searchResults", 
		"date",
	})

	asaiChat := chains.NewLLMChain(llm, personaPrompt)

	input := map[string]interface{}{
		"input":  userInput,
		"history": buffer["history"],
		"searchPerformed": searchPerformed,
		"searchResults": searchResults,
		"date": time.Now().Format("January 02, 2006"),
	}

	answer, err := chains.Predict(ctx, asaiChat, input)
	if err != nil {
		return nil, err
	}

	return map[string]any{"output": answer}, nil
}

func (chain AsaiChain) GetInputKeys() []string {
	return []string{chain.InputKey}
}

func (chain AsaiChain) GetOutputKeys() []string {
	return []string{chain.OutputKey}
}

func (chain AsaiChain) GetMemory() schema.Memory {
	return chain.Memory.Buffer()
}