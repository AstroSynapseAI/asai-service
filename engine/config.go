package engine

import (
	"errors"
	"fmt"

	"github.com/AstroSynapseAI/app-service/engine/agents"
	"github.com/AstroSynapseAI/app-service/engine/agents/email"
	"github.com/AstroSynapseAI/app-service/engine/agents/search"
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

type Config struct {
	DB     *database.Database
	Avatar models.Avatar
}

type AvatarConfig interface {
	LoadConfig(userID uint)
	GetDB() *database.Database
	GetAvatarName() string
	GetAvatarLLM() llms.LanguageModel
	GetAvatarPrimer() string
	GetAvatarMemorySize() int
	AvatarIsPublic() bool
	GetAgents() []tools.Tool
	GetTools() []tools.Tool
}

var _ AvatarConfig = (*Config)(nil)

func NewConfig(db *database.Database) *Config {
	return &Config{
		DB: db,
	}
}

func (cnf *Config) LoadConfig(avatarID uint) {
	avatars := repositories.NewAvatarsRepository(cnf.DB)

	avatar, err := avatars.Fetch(avatarID)
	if err != nil {
		fmt.Println("Error loading avatar:", err)
		return
	}

	cnf.Avatar = avatar
}

func (cnf *Config) GetDB() *database.Database {
	return cnf.DB
}

func (cnf *Config) GetAvatarLLM() llms.LanguageModel {
	avatarLLM := cnf.Avatar.LLM
	activeLLMs := cnf.Avatar.ActiveLLMs
	if len(activeLLMs) == 0 {
		fmt.Println("Avatar has no active LLMs")
		return nil
	}

	// extract active llm where activeLLM.llmID == avatarLLM.ID
	var activeLLM models.ActiveLLM
	for _, active := range activeLLMs {
		if active.LLM.ID == avatarLLM.ID {
			activeLLM = active
		}
	}

	LLM, err := openai.NewChat(
		openai.WithModel(activeLLM.LLM.Slug),
		openai.WithToken(activeLLM.Token),
	)

	if err != nil {
		fmt.Println("Error loading Avatar LLM:", err)
		return nil
	}
	return LLM
}

func (cnf *Config) GetAvatarName() string {
	return cnf.Avatar.Name
}

func (cnf *Config) GetAvatarPrimer() string {
	return cnf.Avatar.Primer
}

func (cnf *Config) GetAvatarMemorySize() int {
	return 4048
}

func (cnf *Config) AvatarIsPublic() bool {
	return cnf.Avatar.IsPublic
}

func (cnf *Config) GetAgents() []tools.Tool {
	activeAgents := cnf.Avatar.ActiveAgents
	loadedAgents := []tools.Tool{}

	for _, activeAgent := range activeAgents {
		agent := agents.NewActiveAgent(cnf.Avatar, activeAgent)

		fmt.Println("Loading agent..." + agent.GetAgentName())
		fmt.Println("agent is active: " + fmt.Sprint(agent.IsAgentActive()))

		if agent.GetAgentSlug() == "search-agent" && agent.IsAgentActive() {
			fmt.Println("Loading search agent...")
			searchAgent, err := search.NewSearchAgent(
				search.WithPrimer(agent.GetAgentPrimer()),
				search.WithLLM(agent.GetAgentLLM()),
				search.WithConfig(agent.GetAgentConfig()),
			)

			if err != nil {
				fmt.Println("Error loading search agent:", err)
				return nil
			}

			loadedAgents = append(loadedAgents, searchAgent)
		}

		if agent.GetAgentSlug() == "email-agent" && agent.IsAgentActive() {
			fmt.Println("Loading email agent...")
			emailAgent, err := email.NewEmailAgent(
				email.WithPrimer(agent.GetAgentPrimer()),
				email.WithLLM(agent.GetAgentLLM().(*openai.Chat)),
				email.WithConfig(agent.GetAgentConfig()),
			)

			if err != nil {
				fmt.Println("Error loading email agent:", err)
				return nil
			}

			loadedAgents = append(loadedAgents, emailAgent)
		}
	}
	return loadedAgents
}

func (cnf *Config) GetTools() []tools.Tool {
	// activeTools := cnf.Avatar.ActiveTools
	var loadedTools []tools.Tool

	return loadedTools
}

func loadActiveLLM(activeLLM models.ActiveLLM) (llms.LanguageModel, error) {

	switch activeLLM.LLM.Slug {
	case "mistral":
		LLM, err := ollama.New(
			ollama.WithModel("mistral"),
			ollama.WithServerURL("http://host.docker.internal:11434/"),
		)

		if err != nil {
			fmt.Println("Error setting mistral:", err)
			return nil, err
		}

		return LLM, nil
	case "gpt-4":
		LLM, err := openai.NewChat(
			openai.WithToken(activeLLM.Token),
			openai.WithModel("gpt-4"),
		)

		if err != nil {
			fmt.Println("Error setting gpt-4:", err)
			return nil, err
		}

		return LLM, nil
	case "gpt-4-turbo-preview":
		LLM, err := openai.NewChat(
			openai.WithToken(activeLLM.Token),
			openai.WithModel("gpt-4-turbo-preview"),
		)

		if err != nil {
			fmt.Println("Error setting gpt-4-turbo-preview:", err)
			return nil, err
		}

		return LLM, nil
	case "gpt-3.5":
		LLM, err := openai.NewChat(
			openai.WithToken(activeLLM.Token),
			openai.WithModel("gpt-3.5"),
		)

		if err != nil {
			fmt.Println("Error setting gpt-3.5:", err)
			return nil, err
		}

		return LLM, nil
	default:
		fmt.Println("Unknown LLM:", activeLLM.LLM.Slug)
		return nil, errors.New("unknown LLM")
	}
}
