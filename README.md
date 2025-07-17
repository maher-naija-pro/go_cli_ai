# go_cli_ai

## Overview

go_cli_ai is an advanced AI-powered command-line interface (CLI) tool written in Go. It works with any OpenAI-compatible endpoint to provide assistant, developer, and tutor functionalities directly from your terminal. The CLI is modular, supports custom prompts, and is easily extensible.

## Features
- Interact with any OpenAI-compatible endpoint from the terminal
- Retrieve and use context directly from your terminal sessions
- Multiple system roles: assistant, developer, tutor
- Customizable prompts via configuration
- Simple command structure using [urfave/cli](https://github.com/urfave/cli)

## Installation

1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd go_cli
   ```
2. **Build the CLI:**
   ```sh
   go build -o ai main.go
   ```

## Configuration

To set up your configuration, run the following command to generate a default `config.yaml` in your current directory:

./ai init-config


```yaml
openai:
  api_key: "sk-..."
  endpoint: "https://api.openai.com/v1/chat/completions"
  model: "gpt-4"
  prompts:
    assistant: "You are a helpful assistant."
    dev: "You are a senior software engineer helping write efficient and secure code."
    tutor: "You are a patient tutor explaining technical topics to beginners."

- **api_key**: Your OpenAI API key
- **endpoint**: Any OpenAI-compatible API endpoint (such as OpenAI or open-source alternatives)
- **model**: Model to use (e.g., gpt-4)
- **prompts**: System prompts for different roles


**Important:** Be sure to replace `"sk-..."` with your actual OpenAI API key in the `api_key` field before using the CLI. Without a valid API key, requests to the OpenAI endpoint will not work.

## Environment Variables

| Variable      | Description                                 | Default        |
|---------------|---------------------------------------------|----------------|
| `CONFIG_PATH` | Path to the configuration YAML file         | `config.yaml`  |

You can set the `CONFIG_PATH` environment variable to specify a custom configuration file location. If not set, the CLI will look for `config.yaml` in the current directory.

## Usage

Run the CLI:
```sh
./ai
```

### Main Commands

- `ask` — Interact with OpenAI endpoint using configured system roles (subcommands: assistant, dev, tutor, etc.)
- `init-config` — Generate a default `config.yaml` file in the current directory

#### Example:
```sh
./ai ask assistant "What can you do?"
./ai ask dev "Write a Go program that reads a file."
./ai ask tutor "Explain Docker to a beginner."
```

##  Integration

The CLI uses streaming responses for real-time output. Prompts and roles are defined in the configuration file. The main logic is in `openai/client.go` and `cmd/openai.go`.

## License

MIT




