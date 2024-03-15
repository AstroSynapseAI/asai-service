package email

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/tools"
)

type TestEmail struct {
	Content string
}

var _ tools.Tool = TestEmail{}

func NewTestEmail() TestEmail {
	return TestEmail{}
}

func (email TestEmail) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Test Email Tool Running...")
	fmt.Println(input)
	return "Email sent...", nil
}

func (email TestEmail) Name() string {
	return "Test Email Agent"
}

func (email TestEmail) Description() string {
	return `
  Email agent expects json of the format as input:
    {
      "content": "Hello World"
    }
  `
}
