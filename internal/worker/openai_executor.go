package worker

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIExecutor struct {
}

func (e *OpenAIExecutor) Execute(ctx context.Context, p NodePayload) (*NodeOutput, error) {
	apiKey, _ := p.InputData["api_key"].(string)
	prompt, _ := p.InputData["prompt"].(string)

	if apiKey == "" {
		return &NodeOutput{Data: map[string]interface{}{"error": "API key missing"}}, nil
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})

	if err != nil {
		return nil, err
	}

	return &NodeOutput{
		Data: map[string]interface{}{"text": resp.Choices[0].Message.Content},
	}, nil
}
