package config

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type OpenAIConfig struct {
	APIKey   string            `yaml:"api_key"`
	Endpoint string            `yaml:"endpoint" `	
	Model    string            `yaml:"model" `
	Prompts  map[string]string `yaml:"prompts"`
}

type Config struct {
	OpenAI OpenAIConfig `yaml:"openai"`
}

// DefaultConfig returns a Config struct with default values.
func DefaultConfig() *Config {
	return &Config{
		OpenAI: OpenAIConfig{
			APIKey:   "",
			Endpoint: "https://api.openai.com/v1/chat/completions",
			Model:    "gpt-4",
				Prompts: map[string]string{
					
					"explain":   "You are a patient tutor who explains technical topics in simple, beginner-friendly language.",
					"dev":       "You are a senior software engineer helping write efficient, secure, and idiomatic code. Offer best practices and clear explanations.",
					"logs":      "You are an expert at interpreting and explaining logs, making them easy for humans to understand and troubleshoot.",
					"debug":     "You are an expert at debugging code, identifying root causes, and suggesting effective fixes.",
					"tests":     "You are an expert at writing comprehensive and maintainable tests for code, including unit and integration tests.",
					"review":    "You are a code reviewer who provides constructive feedback and suggestions for improvement.",

				},
		},
	}
}

func LoadConfig(path string) *Config {
	var cfg Config
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Config file not found at %s: %v. Using default config.", path, err)
		return DefaultConfig()
	}
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		log.Printf("Error unmarshalling config file: %v", err)
		return DefaultConfig()
	}

	// Fill in missing fields with defaults
	if cfg.OpenAI.Endpoint == "" {
		cfg.OpenAI.Endpoint = DefaultConfig().OpenAI.Endpoint
	}
	if cfg.OpenAI.Model == "" {
		cfg.OpenAI.Model = DefaultConfig().OpenAI.Model
	}
	if len(cfg.OpenAI.Prompts) == 0 {
		cfg.OpenAI.Prompts = DefaultConfig().OpenAI.Prompts
	}

	// APIKey is required, so get default from struct if missing
	if cfg.OpenAI.APIKey == "" {
		cfg.OpenAI.APIKey = DefaultConfig().OpenAI.APIKey
	}

	return &cfg
}
