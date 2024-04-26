package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
)

type APITool struct {
	apiChain  chains.Chain
	client    APIClient
	ActiveLLM llms.Model
	APIDocs   string
	APISecret string
	APIKey    string
	APIToken  string
}

type APIClient struct {
	Secret string
	Key    string
	Token  string
}

func (client *APIClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+client.Token)

	return http.DefaultClient.Do(req)
}

func NewTool(options ...ApiOptions) *APITool {
	tool := &APITool{}

	for _, option := range options {
		option(tool)
	}

	tool.apiChain = chains.NewAPIChain(tool.ActiveLLM, http.DefaultClient)

	return tool
}

func (tool *APITool) Name() string {
	return "DNB API Tool"
}

func (tool *APITool) Description() string {
	return "DNB API Tool"
}

func (tool *APITool) Call(ctx context.Context, query string) (string, error) {
	fmt.Println("DNB Agent api tool running...")

	input := map[string]any{
		"api_docs": tool.APIDocs,
		"input":    query,
	}

	result, err := chains.Call(context.Background(), tool.apiChain, input)
	if err != nil {
		log.Println("Api chain failed with error: ", err)
		return "", nil
	}

	response := result["answer"].(string)
	return response, nil
}
