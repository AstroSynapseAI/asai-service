package google

import (
	"context"
	"errors"

	"github.com/tmc/langchaingo/tools"
)

type Tool struct {
	client *Client
}

var _ tools.Tool = &Tool{}

func New(apiKey string, maxResults int) (*Tool, error) {
	return &Tool{
		client: NewClient(apiKey, maxResults),
	}, nil
}

func (tool Tool) Name() string {
	return "Google Search API Tool"
}

func (tool Tool) Description() string {
	return "Google Search API Tool"
}

func (tool Tool) Call(ctx context.Context, input string) (string, error) {
	result, err := tool.client.Search(ctx, input)
	if err != nil {
		if errors.Is(err, ErrNoGoodResult) {
			return "No good Google search results was found", nil
		}
		return "", err
	}

	return result, nil
	
}