package email

import (
	"context"
	"fmt"

	asaiTools "github.com/AstroSynapseAI/app-service/engine/tools"
	"github.com/AstroSynapseAI/app-service/engine/tools/email"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
	"github.com/xhit/go-simple-mail/v2"
)

var _ tools.Tool = &EmailAgent{}

type EmailAgent struct {
	Primer    string
	LLM       llms.LanguageModel
	Executor  agents.Executor
	EmailTool *email.Client
}

func NewEmailAgent(options ...EmailAgentOptions) (*EmailAgent, error) {
	// create a new email agent
	emailAgent := &EmailAgent{
		EmailTool: email.NewClient(
			email.WithHost("mail.gandi.net"),
			email.WithPassword("asai1234"),
			email.WithUsername("dispatch@astrosynapse.com"),
			email.WithEncryption(mail.EncryptionSSLTLS),
			email.WithPort(465),
		),
	}

	// apply email agent options
	for _, option := range options {
		option(emailAgent)
	}

	promptTmplt := prompts.PromptTemplate{
		Template:       emailAgent.Primer,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]interface{}{
			"tool_names":        asaiTools.Names([]tools.Tool{emailAgent.EmailTool}),
			"tool_descriptions": asaiTools.Descriptions([]tools.Tool{emailAgent.EmailTool}),
		},
	}

	agent := agents.NewOneShotAgent(
		emailAgent.LLM,
		[]tools.Tool{emailAgent.EmailTool},
		agents.WithPrompt(promptTmplt),
		agents.WithMaxIterations(3),
		agents.WithMemory(memory.NewSimple()),
	)

	emailAgent.Executor = agents.NewExecutor(agent, []tools.Tool{emailAgent.EmailTool})

	return emailAgent, nil
}

func (emailAgent *EmailAgent) Name() string {
	return "Email Agent"
}

func (emailAgent *EmailAgent) Description() string {
	return `
  Email agent enables sending emails. The agent expects
  email address and email subject and message as input, in the format:
  - send to: email@example.com
  - subject: Email subject
  - message: Email content
  `
}

func (emailAgent *EmailAgent) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Email Agent Running...")
	fmt.Println(input)

	response, err := chains.Run(ctx, emailAgent.Executor, input)
	if err != nil {
		return "Email Agent Error: " + err.Error(), nil
	}

	return response, nil
}
