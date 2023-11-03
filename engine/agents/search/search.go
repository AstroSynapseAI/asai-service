package search

import (
	"context"
	"fmt"
	"os"

	"github.com/AstroSynapseAI/engine/templates"
	"github.com/AstroSynapseAI/engine/tools/google"

	asaiTools "github.com/AstroSynapseAI/engine/tools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/duckduckgo"
)

var _ tools.Tool = &SearchAgent{}

type SearchAgent struct {
	Memory   schema.Memory
	Executor agents.Executor
}

func NewSearchAgent(options ...SearchAgentOptions) (*SearchAgent, error) {
	// create a new search agent
	searchAgent := &SearchAgent{
		Memory: memory.NewSimple(),
	}

	// apply search agent options
	for _, option := range options {
		option(searchAgent)
	}

	// create new llm handle
	llm, err := openai.NewChat(
		openai.WithModel("gpt-4"),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create google search API Tool
	apiKey := os.Getenv("SERPAPI_API_KEY")
	google, err := google.New(apiKey, 10)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create DuckDuckGo search API Tool
	ddg, err := duckduckgo.New(10, duckduckgo.DefaultUserAgent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create search agent tools
	searchTools := []tools.Tool{google, ddg}

	// load custom search agent template
	searchTmplt, err := templates.Load("search.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create search prompt template
	promptTmplt := prompts.PromptTemplate{
		Template:       searchTmplt,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"tool_names":        asaiTools.Names(searchTools),
			"tool_descriptions": asaiTools.Descriptions(searchTools),
			"history":           "",
		},
	}

	// create the search prompt
	agent := agents.NewOneShotAgent(
		llm,
		searchTools,
		agents.WithMemory(searchAgent.Memory),
		agents.WithPrompt(promptTmplt),
		agents.WithMaxIterations(3),
	)

	// create agents executor chain
	searchAgent.Executor = agents.NewExecutor(agent, searchTools)

	return searchAgent, nil
}

func (agent *SearchAgent) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Search Agent called...")
	reponse, err := chains.Run(ctx, agent.Executor, input)
	if err != nil {
		return "Search Agent encountered an error: " + err.Error(), nil
	}
	return reponse, nil
}

func (agent *SearchAgent) Name() string {
	return "Search Agent"
}

func (agent *SearchAgent) Description() string {
	return `
		Search Agent is an agent specialized in searching the web.
		The agent can use DuckDuckGo, SerpApi Google API, and Metaphor Search 
		tools to search the web. Input should be a question or query related to Human input.
		Output should be a summary of the search with the most relevant results.
	`
}
