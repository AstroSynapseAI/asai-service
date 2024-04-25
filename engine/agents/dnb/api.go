package dnb

import (
	"context"
)

type ApiTool struct {
}

func NewApiTool() *ApiTool {
	return &ApiTool{}
}

func (ApiTool) Name() string {
	return "DNB API Tool"
}

func (ApiTool) Description() string {
	return "DNB API Tool"
}

func (ApiTool) Call(ctx context.Context, input string) (string, error) {
	return "", nil
}
