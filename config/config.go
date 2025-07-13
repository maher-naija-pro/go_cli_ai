type Config struct {
	OpenAI struct {
		APIKey   string            `yaml:"api_key"`
		Endpoint string            `yaml:"endpoint"`
		Model    string            `yaml:"model"`
		Prompts  map[string]string `yaml:"prompts"` // Keyed prompts
	} `yaml:"openai"`
}

