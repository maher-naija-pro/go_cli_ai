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
			Model:    "gpt-3.5-turbo",
			Prompts: map[string]string{
				"default": "You are a helpful assistant.",
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
		return nil
	}

	// Fill in missing fields with defaults
	if cfg.OpenAI.Endpoint == "" {
		cfg.OpenAI.Endpoint = "https://api.openai.com/v1/chat/completions"
	}
	if cfg.OpenAI.Model == "" {
		cfg.OpenAI.Model = "gpt-3.5-turbo"
	}
	if cfg.OpenAI.Prompts == nil || len(cfg.OpenAI.Prompts) == 0 {
		cfg.OpenAI.Prompts = map[string]string{
			"default": "You are a helpful assistant.",
		}
	}

	// APIKey is required, so still check for it
	if cfg.OpenAI.APIKey == "" {
		log.Printf("Missing OpenAI API key in config")
		return nil
	}

	return &cfg
}
