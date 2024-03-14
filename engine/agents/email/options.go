package email

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/xhit/go-simple-mail/v2"
)

type EmailAgentOptions func(*EmailAgent)

type config struct {
	IMAPServer string `json:"imap_server"`
	SMTPServer string `json:"smtp_server"`
	IMAPPort   int    `json:"imap_port"`
	SMTPPort   int    `json:"smtp_port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
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

		switch configData.Encryption {
		case "ssl":
			emailAgent.Encryption = mail.EncryptionSSL
		case "tls":
			emailAgent.Encryption = mail.EncryptionTLS
		case "starttls":
			emailAgent.Encryption = mail.EncryptionSTARTTLS
		default:
			emailAgent.Encryption = mail.EncryptionNone
		}

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
