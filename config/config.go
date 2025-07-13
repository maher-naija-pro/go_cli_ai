package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type OpenAIConfig struct {
	APIKey   string            `yaml:"api_key"`
	Endpoint string            `yaml:"endpoint"`
	Model    string            `yaml:"model"`
	Prompts  map[string]string `yaml:"prompts"`
}

type Config struct {
	OpenAI OpenAIConfig `yaml:"openai"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading config file at %s: %v", path, err)
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		log.Printf("Error unmarshalling config file: %v", err)
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Check that all fields are set
	if cfg.OpenAI.APIKey == "" {
		return nil, fmt.Errorf("API key is missing in the config")
	}
	if cfg.OpenAI.Endpoint == "" {
		return nil, fmt.Errorf("Endpoint is missing in the config")
	}
	if cfg.OpenAI.Model == "" {
		return nil, fmt.Errorf("Model is missing in the config")
	}
	if cfg.OpenAI.Prompts == nil || len(cfg.OpenAI.Prompts) == 0 {
		return nil, fmt.Errorf("Prompts are missing in the config")
	}

	log.Printf("Config loaded successfully from %s", path)
	return &cfg, nil
}
