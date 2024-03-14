package email

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/xhit/go-simple-mail/v2"
)

type EmailAgentOptions func(*EmailAgent)

type config struct {
	IMAPServer string
	SMTPServer string
	IMAPPort   int
	SMTPPort   int
	Username   string
	Password   string
	Encryption mail.Encryption
}

func WithLLM(llm *openai.Chat) EmailAgentOptions {
	return func(emailAgent *EmailAgent) {
		emailAgent.LLM = llm
	}
}

func WithConfig(data string) EmailAgentOptions {
	return func(emailAgent *EmailAgent) {
		var configData config

		err := json.Unmarshal([]byte(data), &configData)
		if err != nil {
			fmt.Println("Error decoding config data:", err)
			return
		}

		emailAgent.SMTPServer = configData.SMTPServer
		emailAgent.SMTPPort = configData.SMTPPort
		emailAgent.IMAPServer = configData.IMAPServer
		emailAgent.IMAPPort = configData.IMAPPort
		emailAgent.Username = configData.Username
		emailAgent.Password = configData.Password
		emailAgent.Encryption = configData.Encryption
	}
}

func WithIMAPServer(hostname string) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.IMAPServer = hostname
	}
}

func WithSMTPServer(hostname string) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.SMTPServer = hostname
	}
}

func WithIMAPPort(port int) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.IMAPPort = port
	}
}

func WithSMTPPort(port int) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.SMTPPort = port
	}
}

func WithUsername(username string) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.Username = username
	}
}

func WithPassword(password string) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.Password = password
	}
}

func WithEncryption(encryption mail.Encryption) EmailAgentOptions {
	return func(agent *EmailAgent) {
		agent.Encryption = encryption
	}
}
