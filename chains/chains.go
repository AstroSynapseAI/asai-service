package chains

import (
	"context"

	"github.com/tmc/langchaingo/chains"
)

func RunAsai (ctx context.Context, chain *AsaiChain, input string) (string, error) {
	return chains.Run(ctx, chain, input)
}