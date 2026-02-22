# Golang Minor Projects

A collection of small and intermediate Go projects demonstrating different use cases, libraries, frameworks, and integrations built with the Go programming language. Each project is self-contained in its own directory and illustrates a unique concept — from REST APIs and CRUD applications to serverless deployments and Slack bots.

## Projects

| Project | Description |
|---|---|
| `AWSLambda-tool` | Tools and utility functions for deploying and managing AWS Lambda functions written in Go |
| `bookstore-manager` | A RESTful service for managing a bookstore inventory with full CRUD operations |
| `email-checker` | A CLI utility that checks DNS records (MX, SPF, DMARC) for email domains |
| `go-fiber-CRM` | A Customer Relationship Management API built using the Fiber web framework with SQLite |
| `go-fiber-mongo-hrms` | An HR Management System API built with Fiber and MongoDB |
| `go-server` | A basic Go web server demonstrating static file serving, routing, and form handling |
| `go-serverless-stack` | A serverless application boilerplate built with Go for AWS Lambda deployment |
| `movies-crud` | A RESTful CRUD API for managing movie records using Gorilla Mux and in-memory storage |
| `slack-age-bot` | A Slack bot that calculates and returns a user's age from their year of birth |
| `slack-file-bot` | A Slack bot that uploads files to a designated Slack channel programmatically |

## Prerequisites

Before running any project, ensure you have the following installed:

- [Go](https://go.dev/dl/) 1.18 or higher
- [Git](https://git-scm.com/)
- Any project-specific dependencies (see individual project READMEs)

## Getting Started

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects
```

2. **Navigate to a project folder:**
```bash
   cd movies-crud
```

3. **Install dependencies and run:**
```bash
   go mod tidy
   go run .
```

   Or build and execute:
```bash
   go build -o app
   ./app
```

## Project Structure

Each project directory typically follows this layout:
```
project-name/
├── main.go         # Application entry point
├── go.mod          # Module definition
├── go.sum          # Dependency checksums
└── ...             # Additional source files, configs, or scripts
```

## Contributing

Contributions are welcome! To get started:

1. Fork the repository
2. Create a new branch for your feature or fix
3. Commit your changes with clear messages
4. Push to your fork
5. Open a pull request describing your changes

Please ensure your code follows Go best practices and includes appropriate comments where needed.
