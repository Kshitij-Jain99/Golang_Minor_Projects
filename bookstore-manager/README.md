# Bookstore Manager

A simple command-line and/or web-based Bookstore Management application written in Go. This project demonstrates the use of Go for building backend services, handling HTTP routing, data storage, and common CRUD (Create, Read, Update, Delete) operations for managing books in a store.

## Features

- Add, list, update and delete books
- Search books by title, author or ISBN
- Simple JSON API interface
- Optional persistent storage (e.g., file, SQLite, or in-memory)
- Clean project layout following Go best practices

## Getting Started

### Prerequisites

- Go 1.18+ installed
- Git for cloning the repository

### Installation
```bash
# Clone the repository
git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git

# Change directory to the bookstore manager
cd Golang_Minor_Projects/bookstore-manager

# Build the application
go build -o bookstore-manager
```

### Running Locally
```bash
# Run the application
./bookstore-manager
```

By default, the server listens on a port (e.g., `:8080`). You can configure this via environment variables or configuration file (if implemented).

## API Endpoints

| Method   | Endpoint      | Description          |
|----------|---------------|----------------------|
| GET      | `/books`      | List all books       |
| GET      | `/books/{id}` | Get book by ID       |
| POST     | `/books`      | Create a new book    |
| PUT      | `/books/{id}` | Update existing book |
| DELETE   | `/books/{id}` | Remove a book        |

### Sample Request
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go Programming",
    "author": "Author Name",
    "isbn": "1234567890",
    "price": 399.99
  }'
```

## Project Structure

A typical layout might look like:
```
bookstore-manager/
├── cmd/                # entrypoint(s)
│   └── main.go
├── internal/           # internal packages
│   ├── handlers/       # HTTP handlers
│   ├── models/         # data models
│   └── store/          # storage logic
├── go.mod
└── go.sum
```

## Dependencies

This project may use popular Go packages:

- `net/http` for HTTP server
- `gorilla/mux` or `chi` for routing
- `gorm` or `sqlx` for database (optional)

To install dependencies:
```bash
go mod tidy
```

## Testing

Unit tests and integration tests can be included using Go's testing tools:
```bash
go test ./...
```

## Contribution

Contributions are welcome. Please follow standard GitHub workflow:

1. Fork the repository
2. Create a new branch
3. Add your changes
4. Open a Pull Request
