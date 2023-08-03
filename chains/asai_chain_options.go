package chains

import (
	"github.com/AstroSynapseAI/engine-service/agents"
	"github.com/AstroSynapseAI/engine-service/memory"
)

type ChainOptions func (chain *AsaiChain)

func WithMemory(memory *memory.AsaiMemory) ChainOptions {
	return func(chain *AsaiChain) {
		chain.Memory = memory
	}
}

func WithSearchAgent(agent *agents.SearchAgent) ChainOptions {
	return func(chain *AsaiChain) {
		chain.SearchAgent = agent
	}
}