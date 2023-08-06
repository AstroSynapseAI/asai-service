package agents

import (
	"fmt"
	"os"

	"github.com/AstroSynapseAI/engine-service/templates"
	"github.com/AstroSynapseAI/engine-service/tools"
	"github.com/AstroSynapseAI/engine-service/tools/google"
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
	executor agents.Executor
}

// NewSearchAgent creates a new search agent with the given options.
//
// options: Variadic parameter to customize the search agent.
// Returns a pointer to a SearchAgent and an error if any.
func NewSearchAgent(options ...SearchAgentOptions) (*SearchAgent, error) {
	// create a new search agent
	searchAgent := &SearchAgent{
		memory: memory.NewSimple(),
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
	google, err := google.New(apiKey, google.DefualtMaxResults)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	// create DuckDuckGo search API Tool 
	ddg, err := duckduckgo.New(5, duckduckgo.DefaultUserAgent)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create web scraping tool
	scraper, err := tools.NewScraper()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create search agent tools
	searchTools := []lc_tools.Tool{google, ddg, scraper}

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
		InputVariables: []string{"input", "history", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"tool_names":        tools.Names(searchTools),
			"tool_descriptions": tools.Descriptions(searchTools),
		},
	}

	// create the search prompt
	agent := agents.NewOneShotAgent(
		llm,
		searchTools,
		agents.WithMemory(searchAgent.memory),
		agents.WithPrompt(promptTmplt),
		agents.WithMaxIterations(3),
	)

	// create agents executor chain	
	executor := agents.NewExecutor(agent, searchTools)
	searchAgent.executor = executor

	return searchAgent, nil
}

// Executor returns the executor of the SearchAgent.
//
// This function does not take any parameters.
// It returns an agents.Executor.
func (agent *SearchAgent) Executor() agents.Executor {
	return agent.executor
}

// Name returns the name of the SearchAgent.
//
// It does not take any parameters.
// It returns a string.
func (agent *SearchAgent) Name() string {
	return "Search Agent"	
}

// Description returns the description of the SearchAgent.
//
// It returns a string that describes the SearchAgent and its capabilities.
func (agent *SearchAgent) Description() string {
	return `
		Search Agent is an agent specialized for searching and scarping the web.
		The agent can use DuckDuckGo and SerpApi Google API tools to search the web, 
		and web scraping tool for reading and scraping various valid urls.
	`
}
