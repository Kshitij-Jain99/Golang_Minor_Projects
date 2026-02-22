# go-serverless-stack

A simple serverless backend written in Go intended for deployment using the Serverless Framework with AWS Lambda and API Gateway. This project demonstrates how to organize Go code, build serverless functions, and integrate with cloud provider services.

## Overview

This project provides a basic boilerplate for:

- Writing serverless handlers in Go
- Structuring a Serverless service directory
- Building and packaging Go functions for cloud deployment
- Handling HTTP events (API Gateway) in a serverless context

## Requirements

- Go 1.18 or higher
- [Serverless Framework](https://www.serverless.com/) (`npm install -g serverless`)
- An AWS account (or equivalent cloud provider)
- AWS credentials configured locally (`aws configure`)

## Setup

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/go-serverless-stack
```

2. **Install Go dependencies:**
```bash
   go mod download
```

3. **Build your Go functions for Linux deployment:**
```bash
   GOOS=linux GOARCH=amd64 go build -o bin/handler main.go
```

   > Adjust the output path and source file to match your handler names as needed.

## Serverless Configuration

Deployments are configured via `serverless.yml`, which defines the provider, functions, and HTTP event triggers. Example configuration:
```yaml
service: go-serverless-stack

provider:
  name: aws
  runtime: go1.x
  region: us-east-1

functions:
  hello:
    handler: bin/handler
    events:
      - http:
          path: hello
          method: get
```

Adjust function names, paths, and triggers based on your implementation.

## Deployment

Deploy to your cloud provider using the Serverless Framework:
```bash
serverless deploy
```

This will:

- Package your compiled Go binaries
- Provision cloud functions (e.g., AWS Lambda)
- Set up API Gateway routes and triggers
- Output the live endpoint URLs on completion

## Local Testing

Run handlers locally before deploying:
```bash
go run main.go
```

Or emulate the serverless environment using the offline plugin:
```bash
serverless offline
```

## Example Handler

A minimal serverless handler in Go:
```go
package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (string, error) {
    return fmt.Sprintf("Hello from serverless Go!"), nil
}

func main() {
    lambda.Start(Handler)
}
```

This function returns a greeting string when triggered by an event. Extend the handler logic to suit your use case.

## How It Works

In a serverless architecture:

- Functions are **stateless** and triggered by events (e.g., HTTP requests via API Gateway)
- The Serverless Framework packages your code and provisions cloud infrastructure
- AWS Lambda runs your Go binary on demand with automatic scaling
- You are billed only for actual execution time and resources consumed

Go's performance and low startup latency make it well-suited for Lambda deployments.

## Project Structure
```
go-serverless-stack/
├── handler/            # Serverless function handlers
├── bin/                # Compiled function binaries (generated at build time)
├── serverless.yml      # Serverless Framework deployment config
├── main.go             # Application entry point
├── go.mod
└── go.sum
```

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/aws/aws-lambda-go/lambda` | AWS Lambda function handler runtime |

All dependencies are managed via Go modules.
