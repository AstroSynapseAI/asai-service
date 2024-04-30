package api

import (
	"context"

	"github.com/struki84/dnbclient"
)

type SearchClient struct {
	client *dnbclient.Client
}

func NewSearch() *SearchClient {
	searchClient := &SearchClient{}

	c, err := dnbclient.NewClient(
		dnbclient.WithAPIKey(""),
	)

	if err != nil {
		return nil

	}

	searchClient.client = c

	return searchClient
}

func (client *SearchClient) Name() string {
	return ""
}

func (client *SearchClient) Description() string {
	return ""
}

func (client *SearchClient) Call(ctx context.Context, input string) (string, error) {
	return "", nil
}
