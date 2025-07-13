package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenAI struct {
		APIKey   string            `yaml:"api_key"`
		Endpoint string            `yaml:"endpoint"`
		Model    string            `yaml:"model"`
		Prompts  map[string]string `yaml:"prompts"`
	} `yaml:"openai"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(raw, &cfg)
	return &cfg, err
}

