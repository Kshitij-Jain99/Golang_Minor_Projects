# slack-file-bot

A simple Slack file upload bot written in Go that uploads specified files to a designated Slack channel using the official Slack Go SDK. The bot supports uploading files programmatically with minimal configuration.

## Features

- Uploads one or more files to a Slack channel using a bot token
- Built with the official [`slack-go/slack`](https://github.com/slack-go/slack) SDK
- Easy configuration via environment variables

## Requirements

- Go 1.18 or later
- A Slack workspace with a bot user configured
- Bot permissions to upload files to channels (`files:write` scope)
- Valid `SLACK_BOT_TOKEN` and target `CHANNEL_ID`

## Installation

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/slack-file-bot
```

2. **Install dependencies:**
```bash
   go mod download
```

3. **Set environment variables:**
```bash
   export SLACK_BOT_TOKEN=<your-slack-bot-token>
   export CHANNEL_ID=<target-slack-channel-id>
```

   > Ensure the bot user has been invited to the target channel in Slack.

4. **Prepare files for upload:**

   Place the files you want to upload in the project directory, or update the `files` slice in `main.go` with the correct file paths.

## Usage

Run the bot:
```bash
go run main.go
```

The bot will iterate over the configured file list and upload each one to the specified Slack channel. Successful uploads and any errors are printed to the console.

## How It Works

On execution, the bot:

1. Reads `SLACK_BOT_TOKEN` and `CHANNEL_ID` from environment variables
2. Initializes a Slack API client via `slack.New()`
3. Iterates over the list of file paths
4. Opens and uploads each file to Slack using `UploadFileV2Context`

### Core Upload Logic
```go
files := []string{"pplx-at-work.pdf"}
for _, f := range files {
    fileReader, err := os.Open(f)
    // error handling ...
    _, err = api.UploadFileV2Context(context.Background(), slack.UploadFileV2Parameters{
        Filename: f,
        Channel:  channel,
        Reader:   fileReader,
    })
    // log results ...
}
```

## Project Structure
```
slack-file-bot/
├── main.go     # Bot setup and file upload logic
├── go.mod      # Module definition and dependencies
└── go.sum      # Dependency checksums
```

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/slack-go/slack` | Official Slack client library for Go |

## Notes

- **Security:** Never hard-code tokens in source code. Always use environment variables or a secure secrets manager.
- Ensure your Slack app has the `files:write` OAuth scope enabled.
- The bot must be invited to the target channel before it can upload files.
