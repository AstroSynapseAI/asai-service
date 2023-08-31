package browser

import (
	"context"
	"strings"

	"github.com/AstroSynapseAI/engine-service/tools/scraper"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/tools"
)

var _ tools.Tool = &BrowserAgent{}

const (
	
	ErrScraping 		= "Browser Agent failed to scrape web"
	ErrLoadingDocuments = "Browser Agent failed to load web into documents"
)

type BrowserAgent struct {
	llm 	*openai.LLM
	Scraper *scraper.Scraper
	Memory 	schema.Memory
} 

func New(options ...BrowserAgentOptions) (*BrowserAgent, error) {
	var err error
	browserAgent := &BrowserAgent{
		Memory: memory.NewSimple(),
	}

	for _, option := range options {
		option(browserAgent)
	}

	browserAgent.llm, err = openai.New(openai.WithModel("gpt-4"))
	if err != nil {
		return nil, err
	}

	browserAgent.Scraper, err = scraper.New()
	if err != nil {
		return nil, err
	}
	
	return browserAgent, nil
}

func (agent *BrowserAgent) Call(ctx context.Context, input string) (string, error) {
	webContent, err := agent.Scraper.Call(ctx, input)
	if err != nil {
		return ErrScraping, nil
	}

	webDocuments, err := agent.loadWebContent(ctx, webContent)
	if err != nil {
		return ErrLoadingDocuments, nil
	}

	summaryChain := chains.LoadStuffSummarization(agent.llm)
	summary, err := chains.Call(
		ctx,
		summaryChain,
		map[string]any{"input_documents": webDocuments},
	)

	response := summary["text"].(string)
	return response, nil
} 

func (agent *BrowserAgent) Name() string {
	return "Web Browser Agent"	
}

func (agent *BrowserAgent) Description() string {
	return `
		Web Browser Agent is an agent specialized in scraping and reading the web pages.
	`
}

func (agent *BrowserAgent) loadWebContent(ctx context.Context, input string) ([]schema.Document, error) {
	webContent, err := agent.Scraper.Call(ctx, input)
	if err != nil {
		return []schema.Document{}, err
	}

	webContentReader := strings.NewReader(webContent)

	loader := documentloaders.NewText(webContentReader)
	if err != nil {
		return []schema.Document{}, err
	}

	spliter := textsplitter.NewTokenSplitter()
	spliter.ChunkSize = 7500
	spliter.ChunkOverlap = 1024
	spliter.ModelName = "gpt-4"

	webDocuments, err := loader.LoadAndSplit(ctx, spliter)
	if err != nil {
		return []schema.Document{}, err
	}

	return webDocuments, nil
}