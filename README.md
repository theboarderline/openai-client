## Golang Client for ChatGPT
Open AI docs: https://platform.openai.com/docs/guides/chat.

### Install

`go get github.com/theboarderline/openai-client`

### Usage

```go
package main

import (
	"github.com/theboarderline/openai-client/turbo"
)

func main() {
    c, _ := turbo.NewClient("sk-xxxxxxxxxxxxxxxxxxxx", "org-xxxxxxxxxxxxxxxxxxxx")
    req := &turbo.Request{
        Model: turbo.ModelGpt35Turbo,
		Messages: []*turbo.Message{
			{
				Role:    turbo.RoleUser,
				Content: "Hello",
			},
		},
	}

	res, err := c.GetChat(req)
	if err != nil {
		panic(err)
	}

	println(res.Choices[0].Message.Content)
	println(res.Usage.PromptTokens)
	println(res.Usage.CompletionTokens)
	println(res.Usage.TotalTokens)
}

```
