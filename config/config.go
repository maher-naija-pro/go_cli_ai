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
	log.Printf("Config loaded successfully from %s", path)
	return &cfg, nil
}
