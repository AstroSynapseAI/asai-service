package agents

import (
	"context"
	"fmt"

	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/AstroSynapseAI/engine-service/tools"
	a "github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	t "github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/duckduckgo"
)

type SearchAgent struct {
	Memory schema.Memory	
}

func NewSearchAgent(options ...SearchAgentOptions) *SearchAgent {
	return applySearchOptions()
}

func (agent *SearchAgent)Prompt(input string) string {
	llm, err := openai.NewChat(
		openai.WithModel("gpt-4"),
	)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	ddg, err := duckduckgo.New(5, "")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	searchTools := []t.Tool{ddg}

	searchTmplt, err := templates.Load("search.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	promptTmplt := prompts.PromptTemplate{
		Template:       searchTmplt,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "history", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"tool_names":        tools.Names(searchTools),
			"tool_descriptions": tools.Descriptions(searchTools),
		},
	}

	executor, err := a.Initialize(
		llm,
		searchTools,
		a.ZeroShotReactDescription,
		a.WithMemory(agent.Memory),
		a.WithPrompt(promptTmplt),
		a.WithMaxIterations(3),
	)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	answer, err := chains.Run(context.Background(), executor, input)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	
	return answer
}

