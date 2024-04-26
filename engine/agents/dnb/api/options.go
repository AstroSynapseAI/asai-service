package api

import "github.com/tmc/langchaingo/llms"

type ApiOptions func(*APITool)

func WithActiveLLM(llm llms.Model) ApiOptions {
	return func(apiTool *APITool) {
		apiTool.ActiveLLM = llm
	}
}

func WithApiDocs(data string) ApiOptions {
	return func(apiTool *APITool) {
		apiTool.APIDocs = data
	}
}
