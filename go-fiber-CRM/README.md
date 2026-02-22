# go-fiber-CRM

A simple RESTful CRM API built in Go using the Fiber web framework and SQLite database. This project provides basic lead management endpoints for creating, retrieving, and deleting leads.

## Features

- REST API for managing lead records
- Built with Go and [Fiber](https://gofiber.io/) (an Express-inspired web framework)
- Uses SQLite as a lightweight embedded database
- Basic CRUD operations for leads

## Requirements

- Go 1.18 or higher
- No external database server required (uses SQLite)

## Installation

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/go-fiber-CRM
```

2. **Download dependencies:**
```bash
   go mod download
```

3. **Build and run:**
```bash
   go run main.go
```

   The server will start on **port 3000** by default.

## API Endpoints

### Get all leads
```
GET /api/v1/lead
```
Returns a list of all leads in the database.

---

### Get a specific lead
```
GET /api/v1/lead/:id
```
Returns the lead with the specified ID.

---

### Create a new lead
```
POST /api/v1/lead
```
Request body (JSON):
```json
{
  "name": "John Doe",
  "company": "Acme Inc",
  "email": "john@example.com",
  "phone": 1234567890
}
```

---

### Delete a lead
```
DELETE /api/v1/lead/:id
```
Removes the lead with the specified ID from the database.

## Database

This project uses a local SQLite database file named `leads.db`. On startup, the application:

- Connects to the SQLite database
- Auto-migrates the `Lead` model (creates the table if it doesn't exist)

## Project Structure
```
go-fiber-CRM/
├── database/       # Database connection setup
├── lead/           # Lead model and route handlers
├── main.go         # Application entry point
├── go.mod
└── go.sum
```

## How It Works

The application initializes a Fiber server, connects to a SQLite database via GORM, and registers routes for lead operations. Fiber context handlers are used to parse incoming requests and return JSON responses.
