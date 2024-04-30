package api

import "github.com/tmc/langchaingo/llms"

type ClientOptions func(*Client)

func WithActiveLLM(llm llms.Model) ClientOptions {
	return func(apiClient *Client) {
		apiClient.ActiveLLM = llm
	}
}

func WithApiDocs(data string) ClientOptions {
	return func(apiClent *Client) {
		apiClent.APIDocs = data
	}
}

func WithAPIToken(token string) ClientOptions {
	return func(apiClient *Client) {
		apiClient.APIToken = token
	}
}
