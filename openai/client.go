package openai

import (
	"bytes"
	"encoding/json"
	"ai/config"
	"io"
	"net/http"
)

func Ask(cfg *config.Config, systemPrompt, userPrompt string) (string, error) {
	reqBody := Request{
		Model: cfg.OpenAI.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", cfg.OpenAI.Endpoint, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+cfg.OpenAI.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	return string(resBody), nil
}

