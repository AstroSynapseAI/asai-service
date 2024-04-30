package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
)

type Client struct {
	apiChain  chains.Chain
	client    APIClient
	ActiveLLM llms.Model
	APIDocs   string
	APIToken  string
}

type APIClient struct {
	Token string
}

func (client *APIClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+client.Token)

	return http.DefaultClient.Do(req)
}

func NewClient(options ...ClientOptions) *Client {
	tool := &Client{}

	for _, option := range options {
		option(tool)
	}

	HTTPAPIClient := &APIClient{Token: tool.APIToken}

	tool.apiChain = chains.NewAPIChain(tool.ActiveLLM, HTTPAPIClient)

	return tool
}

func (tool *Client) Name() string {
	return "DNB API Tool"
}

func (tool *Client) Description() string {
	return "DNB API Tool"
}

func (tool *Client) Call(ctx context.Context, query string) (string, error) {
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
