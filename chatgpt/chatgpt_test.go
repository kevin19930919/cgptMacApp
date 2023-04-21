package chatgpt

import (
	"strings"
	"testing"
)

func TestSendSingleRequest(t *testing.T) {
	apiKey := "your_api_key_here"
	c := &ChatGPTFunc{APIKEY: apiKey}

	messages := []ChatGPTMessage{
		{Role: "system", Content: "Your test instruction here."},
		{Role: "user", Content: "Your test input here."},
	}

	args := ChatGPTArgs{Message: messages}

	result, err := c.sendSingleRequest(args)
	if err != nil {
		t.Errorf("SendSingleRequest() returned an error: %v", err)
	}

	if result == nil {
		t.Error("SendSingleRequest() returned a nil result")
	} else {
		t.Logf("SendSingleRequest() returned a response: %s", *result)
		if strings.TrimSpace(*result) == "" {
			t.Error("SendSingleRequest() returned an empty response")
		}
	}
}
