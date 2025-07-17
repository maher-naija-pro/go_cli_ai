package config

import (
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

func LoadConfig(path string) (*Config) {
	var cfg Config
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading config file at %s: %v", path, err)
		return nil
	}
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		log.Printf("Error unmarshalling config file: %v", err)
		return nil
	}

	// Check that all fields are set
	if cfg.OpenAI.APIKey == "" {
		log.Printf("Missing OpenAI API key in config")
		return nil 
	}
	if cfg.OpenAI.Endpoint == "" {
		log.Printf("Missing OpenAI endpoint in config")
		return nil 
	}
	if cfg.OpenAI.Model == "" {
		log.Printf("Missing OpenAI model in config")
		return nil
	}
	if cfg.OpenAI.Prompts == nil || len(cfg.OpenAI.Prompts) == 0 {
		log.Printf("Missing OpenAI prompts in config")
		return nil
	}

	return &cfg
}
