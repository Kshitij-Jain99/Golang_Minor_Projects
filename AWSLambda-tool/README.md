# AWSLambda-tool

A command-line tool written in Go to interact with and manage AWS Lambda functions. This tool simplifies common Lambda operations such as listing functions, invoking them, and retrieving logs — all from your terminal.

## Features

- List all AWS Lambda functions in a region
- Invoke Lambda functions with custom payloads
- Retrieve and display function logs
- View function configuration and metadata
- Lightweight and fast, powered by Go

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher
- An AWS account with appropriate IAM permissions
- AWS credentials configured locally via:
  - `~/.aws/credentials` file, or
  - Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`)

## Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
cd Golang_Minor_Projects/AWSLambda-tool
```

Install dependencies:

```bash
go mod tidy
```

Build the binary:

```bash
go build -o awslambda-tool .
```

## Usage

### List Lambda Functions

```bash
./awslambda-tool list
```

### Invoke a Lambda Function

```bash
./awslambda-tool invoke --function <function-name> --payload '{"key": "value"}'
```

### Get Function Details

```bash
./awslambda-tool describe --function <function-name>
```

## Configuration

You can specify the AWS region using the `--region` flag or by setting the `AWS_REGION` environment variable:

```bash
./awslambda-tool list --region us-east-1
```

## Project Structure

```
AWSLambda-tool/
├── main.go         # Entry point
├── go.mod          # Go module definition
├── go.sum          # Dependency checksums
└── README.md       # Project documentation
```

## Dependencies

- [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2) — Official AWS SDK for interacting with Lambda and other AWS services

## IAM Permissions

Ensure the IAM user or role you're using has the following permissions:

```json
{
  "Effect": "Allow",
  "Action": [
    "lambda:ListFunctions",
    "lambda:InvokeFunction",
    "lambda:GetFunction",
    "logs:GetLogEvents"
  ],
  "Resource": "*"
}
```
