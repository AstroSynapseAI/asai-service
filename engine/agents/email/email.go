package email

import (
	"context"
	"fmt"

	"github.com/AstroSynapseAI/app-service/engine/tools/email"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
	"github.com/xhit/go-simple-mail/v2"

	"github.com/tmc/langchaingo/llms/openai"
)

var _ tools.Tool = &EmailAgent{}

const (
	JsonPrompt = `
  You are a helpful assistant designed to output JSON only.

  You will receive the following as input:
  - send to: email@example.com
  - subject: Email subject
  - message: Email content

  Your output should be in JSON format only.
  {
    "sendTo": "email@example.com",
    "subject": "Email subject",
    "message": "Email HTML content"
  }
  `
)

type EmailAgent struct {
	Primer     string
	LLM        *openai.Chat
	Executor   agents.Executor
	EmailTool  *email.Client
	IMAPServer string
	SMTPServer string
	IMAPPort   int
	SMTPPort   int
	Password   string
	Username   string
	Encryption mail.Encryption
}

func NewEmailAgent(options ...EmailAgentOptions) (*EmailAgent, error) {
	// create a new email agent

	emailAgent := &EmailAgent{}

	// apply email agent options
	for _, option := range options {
		option(emailAgent)
	}

	// fmt.Println("Host: ", emailAgent.SMTPServer)
	// fmt.Println("Port: ", emailAgent.SMTPPort)
	// fmt.Println("Username: ", emailAgent.Username)
	// fmt.Println("Password: ", emailAgent.Password)
	// fmt.Println("Encryption: ", emailAgent.Encryption)

	emailClient := email.NewClient(
		email.WithHost(emailAgent.SMTPServer),
		email.WithPassword(emailAgent.Password),
		email.WithUsername(emailAgent.Username),
		email.WithEncryption(emailAgent.Encryption),
		email.WithPort(emailAgent.SMTPPort),
	)

	emailAgent.EmailTool = emailClient

	return emailAgent, nil
}

func (emailAgent *EmailAgent) Name() string {
	return "Email Agent"
}

func (emailAgent *EmailAgent) Description() string {
	return `
  Email agent enables sending emails. The agent expects
  email address, email subject and the email message as input. 
  Please use the following format:
  
  - send to: email@example.com
  - subject: Email subject
  - message: Email content
  `
}

func (emailAgent *EmailAgent) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Email Agent Running...")
	fmt.Println(input)

	msg := []schema.ChatMessage{
		schema.SystemChatMessage{Content: JsonPrompt},
		schema.HumanChatMessage{Content: input},
	}

	response, err := emailAgent.LLM.Call(ctx, msg)
	if err != nil {
		return "Email Agent Error: " + err.Error(), nil
	}

	jsonResponse := response.GetContent()

	toolResponse, err := emailAgent.EmailTool.Call(ctx, jsonResponse)
	if err != nil {
		return "Email Agent Error: " + err.Error(), nil
	}

	return toolResponse, nil
}
