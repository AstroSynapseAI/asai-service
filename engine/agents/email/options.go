package email

import (
	"github.com/tmc/langchaingo/llms/openai"
)

type EmailAgentOptions func(*EmailAgent)

func WithDefaultPrimer() EmailAgentOptions {
	return func(emailAgent *EmailAgent) {
		emailAgent.Primer = defaulPrimer()
	}
}

func WithLLM(llm *openai.Chat) EmailAgentOptions {
	return func(emailAgent *EmailAgent) {
		emailAgent.LLM = llm
	}
}

func defaulPrimer() string {
	return `
  Email agent is trained to write and send emails using the following tool:

  {{.tool_descriptions}}
  
  The agent MUST provide the correct json input for the Email tool.
 
  To use the tool, you MUST the following format:

  Thought: Do I need to use the Email tool? Yes
  Action: the action to take, should be {{.tool_names}}
  Action Input: the input to the action
  Observation: the result of the action
  Final Answer: report the result of the action 

  If you you do not receive all 3 input fields, you MUST respond with the following format:

  Thought: Do I have all the required information to use the Email tool? No
  Final Answer: Please provide all the required information to use the Email tool


  Your final answer MUST have the prefix "Final Answer:"

  Begin!

  New Input: {{.input}}

  Thought:{{.agent_scratchpad}}
  
  `
}
