# go-fiber-mongo-hrms

A basic Human Resource Management System (HRMS) REST API built in Go using the Fiber web framework and MongoDB as the database backend. This API provides endpoints for managing employee/HR records.

## Features

- RESTful API built with [Fiber](https://gofiber.io/)
- Data persistence with MongoDB
- Full CRUD operations for employee/HR records

## Requirements

- Go 1.18 or higher
- MongoDB (local or cloud instance)

## Setup

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/go-fiber-mongo-hrms
```

2. **Install dependencies:**
```bash
   go mod download
```

3. **Configure environment variables:**

   Create a `.env` file in the project root and define your MongoDB connection details:
```env
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=hrms
   COLLECTION_NAME=employees
```

4. **Run the application:**
```bash
   go run main.go
```

   The server will start on **port 3000** by default.

## API Endpoints

### Get all employees
```
GET /employee
```
Returns a list of all employee records stored in the database.

---

### Create a new employee
```
POST /employee
```
Request body (JSON):
```json
{
  "name": "Alice Smith",
  "email": "alice@example.com",
  "position": "Developer",
  "salary": 60000
}
```

---

### Get a specific employee
```
GET /employee/:id
```
Returns the employee record with the specified ID.

---

### Update an employee
```
PUT /employee/:id
```
Updates fields for an existing employee record.

---

### Delete an employee
```
DELETE /employee/:id
```
Removes the employee with the specified ID from the database.

> **Note:** Exact route names and field definitions may vary depending on the implementation.

## Database

This application uses MongoDB for data persistence. Ensure your MongoDB instance is running and accessible before starting the application. Update the connection configuration in `.env` as needed.

## How It Works

On startup, the application:

1. Connects to the MongoDB database using the provided URI
2. Registers HTTP routes via Fiber
3. Handles CRUD operations through controller functions that interact with the MongoDB collection

## Project Structure
```
go-fiber-mongo-hrms/
├── controllers/    # HTTP handler logic
├── database/       # Database connection setup
├── routes/         # Route definitions
├── models/         # Data models (MongoDB document schemas)
├── main.go         # Application entry point
├── go.mod
└── go.sum
```

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/gofiber/fiber/v2` | Fiber web framework |
| `go.mongodb.org/mongo-driver/mongo` | Official MongoDB driver for Go |
