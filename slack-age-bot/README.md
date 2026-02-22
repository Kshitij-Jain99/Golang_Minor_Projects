# slack-age-bot

A simple Slack bot written in Go that calculates a user's age based on their year of birth. The bot listens for commands in a Slack workspace using the Slacker library and responds with the computed age.

## Features

- Listens for a natural language command to calculate age
- Built with Go and the [Slacker](https://github.com/shomali11/slacker) package
- Connects to Slack via the RTM API using bot and app tokens
- Prints command analytics to the console

## Requirements

- Go 1.18 or higher
- A Slack workspace
- A Slack app with:
  - Bot token (`SLACK_BOT_TOKEN`)
  - App token (`SLACK_APP_TOKEN`)
  - Required scopes added and app installed in your workspace

## Installation

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/slack-age-bot
```

2. **Install dependencies:**
```bash
   go mod download
```

3. **Set environment variables:**

   Export your Slack tokens in your shell or add them to a `.env` file:
```bash
   export SLACK_BOT_TOKEN=<your_bot_token>
   export SLACK_APP_TOKEN=<your_app_token>
```

   > **Security note:** Never hard-code tokens in source code. Use environment variables or a secrets manager in production.

## Usage

Start the bot:
```bash
go run main.go
```

Invite the bot to a channel in your Slack workspace, then send the command:
```
my yob is 1990
```

The bot will reply with the calculated age based on the current year.

## Example

**Command sent in Slack:**
```
my yob is 1995
```

**Bot response:**
```
age is 30
```

*(actual result depends on the current year)*

## How It Works

On startup, the bot:

1. Connects to the Slack RTM API using your bot and app tokens
2. Registers a command handler matching the pattern `"my yob is <year>"`
3. Parses the year parameter from the incoming message
4. Calculates age as `current year - year of birth`
5. Sends the computed age back as a reply in Slack

## Project Structure
```
slack-age-bot/
├── main.go     # Bot setup, command registration, and age logic
├── go.mod      # Module definition and dependencies
└── go.sum      # Dependency checksums
```

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/shomali11/slacker` | Slack bot framework for Go |
| `github.com/slack-go/slack` | Underlying Slack API client (via Slacker) |
| `github.com/gorilla/websocket` | WebSocket support for RTM connection |

All dependencies are managed via Go modules.

## Notes

- Ensure your Slack app has the appropriate OAuth scopes to receive and reply to messages.
- Tokens should be stored securely and never committed to version control.
