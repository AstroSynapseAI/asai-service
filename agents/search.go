package agents

import (
	"fmt"

	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/AstroSynapseAI/engine-service/tools"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools/duckduckgo"

	lc_tools "github.com/tmc/langchaingo/tools"
)

type SearchAgent struct {
	memory   schema.Memory
	context  any	
	executor agents.Executor
}

func NewSearchAgent(options ...SearchAgentOptions) (*SearchAgent, error) {
	// create a new search agent
	searchAgent := &SearchAgent{
		memory: memory.NewSimple(),
	}

	// apply search agent options
	for _, option := range options {
		option(searchAgent)
	}

	llm, err := openai.NewChat(
		openai.WithModel("gpt-4"),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ddg, err := duckduckgo.New(5, "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	searchTools := []lc_tools.Tool{ddg}

	searchTmplt, err := templates.Load("search.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
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

	agent := agents.NewOneShotAgent(
		llm,
		searchTools,
		agents.WithMemory(searchAgent.memory),
		agents.WithPrompt(promptTmplt),
		agents.WithMaxIterations(3),
	)

	executor := agents.NewExecutor(agent, searchTools)

	searchAgent.executor = executor

	return searchAgent, nil
}

func (agent *SearchAgent) Executor() agents.Executor {
	return agent.executor
}

func (agent *SearchAgent) Name() string {
	return "Search Agent"	
}

func (agent *SearchAgent) Description() string {
	return `
		Search Agent is an agent specialized for searching and scarping the web.
		The agent can use DuckDuckGo and SerpApi Google API tools to search the web, 
		and web scraping tool for reading and scraping various valid urls.
	`
}
