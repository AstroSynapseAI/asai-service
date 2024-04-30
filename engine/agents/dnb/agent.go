package dnb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AstroSynapseAI/app-service/engine/agents/dnb/api"
	util "github.com/AstroSynapseAI/app-service/engine/tools"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

type DNBAgent struct {
	Primer   string
	LLM      llms.Model
	Executor *agents.Executor
	Config   config
	Tools    []tools.Tool
}

var _ tools.Tool = &DNBAgent{}

func NewDNBAgent(options ...DNBAgentOptions) (*DNBAgent, error) {
	dnbAgent := &DNBAgent{}

	for _, option := range options {
		option(dnbAgent)
	}

	apiTool := api.NewTool(
		api.WithActiveLLM(dnbAgent.LLM),
		api.WithApiDocs(dnbAgent.loadAPIDocs(""))
	)

	dnbAgent.Tools = []tools.Tool{
		NewDocummentTool(),
		apiTool,
	}

	agent := agents.NewOneShotAgent(
		dnbAgent.LLM,
		dnbAgent.Tools,
		agents.WithPrompt(dnbAgent.loadTemplate()),
		agents.WithMaxIterations(5),
	)

	dnbAgent.Executor = agents.NewExecutor(agent, dnbAgent.Tools)

	return dnbAgent, nil
}

func (dnbAgent DNBAgent) Name() string {
	return "DNB"
}

func (dnbAgent DNBAgent) Description() string {
	return "DNB agent"
}

func (dnbAgent DNBAgent) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("DNB Agent Running...")

	response, err := chains.Run(ctx, dnbAgent.Executor, input)
	if err != nil {
		return "DNB Agent encountered an error: " + err.Error(), nil
	}

	return response, nil
}

func (dnbAgent *DNBAgent) loadTemplate() prompts.PromptTemplate {
	// load template
	return prompts.PromptTemplate{
		Template:       dnbAgent.Primer,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"today":             "2023-01-01",
			"tool_names":        util.Names(dnbAgent.Tools),
			"tool_descriptions": util.Descriptions(dnbAgent.Tools),
			"history":           "",
		},
	}
}

func (dnbAgent *DNBAgent) loadAPIDocs(file string) string {
	path := ""
	docs, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading api docs: ", err)
	}

	return string(docs)
}
