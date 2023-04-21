package chatgpt

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	MODEL   = "gpt-3.5-turbo"
	BASEURL = "https://api.openai.com/v1"
)

type TurboUserType string

const (
	TurboSystem = TurboUserType("system")
	// TODO: for the use case right now, we don't need assistant yet
	TurboUser = TurboUserType("user")
)

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Message struct {
	Role    TurboUserType `json:"role"`
	Content string        `json:"content"`
}

type CompletionChoice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   Usage              `json:"usage"`
}

type ChatGPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTArgs struct {
	Message []ChatGPTMessage `json:"messages"`
}

type ChatGPTFunc struct {
	sync.Mutex
	APIKEY string
	Result string
}

func (c *ChatGPTFunc) SendRequest(messages ChatGPTArgs) *string {
	c.Lock()
	defer c.Unlock()

	var wg sync.WaitGroup
	results := make(chan *string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := c.sendSingleRequest(messages)
			results <- result
		}()
	}

	wg.Wait()
	close(results)

	finalResult := constructFinalResult(results)
	return finalResult
}

func (c *ChatGPTFunc) sendSingleRequest(messages ChatGPTArgs) *string {
	// Set up the API request headers and data
	data := map[string]interface{}{
		"model":    MODEL,
		"messages": messages.Message,
	}

	// Make the API request
	client := resty.New().
		SetTimeout(30*time.Second).
		SetRetryCount(1).
		SetHeader("content-type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.APIKEY)).
		SetBaseURL(BASEURL)

	req := client.R().SetBody(data)
	resp, err := req.Post("/chat/completions")

	// Handle the API response
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	} else {
		var resBody CompletionResponse
		if err := json.Unmarshal(resp.Body(), &resBody); err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		fmt.Println("successfully get response", resp)
		return &resBody.Choices[0].Message.Content
	}
}

func constructFinalResult(results chan *string) *string {
	var finalResult string
	for result := range results {
		if result != nil {
			if len(finalResult) > 0 {
				// to make response more readable
				finalResult += "\n\n"
			}
			finalResult += *result
		}
	}
	return &finalResult
}
