# movies-crud

A simple RESTful CRUD API for managing movies, built with Go using the Gorilla Mux router. This API demonstrates basic Create, Read, Update, and Delete operations on an in-memory list of movies.

## Features

- Full CRUD REST API for movie management
- In-memory data store — no database required
- Routing powered by [Gorilla Mux](https://github.com/gorilla/mux)

## Requirements

- Go 1.18 or higher
- `github.com/gorilla/mux`

## Installation

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/movies-crud
```

2. **Fetch dependencies:**
```bash
   go mod download
```

3. **Run the server:**
```bash
   go run main.go
```

   The server starts listening on **port 8000** by default.

## API Endpoints

All endpoints operate under the `/movies` path.

### Get all movies
```
GET /movies
```
Returns a list of all movies in JSON format.

---

### Get a single movie
```
GET /movies/{id}
```
Returns the movie matching the provided `id`.

---

### Create a new movie
```
POST /movies
```
Request body (JSON):
```json
{
  "isbn": "123456789",
  "title": "New Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Smith"
  }
}
```

A new movie is added with an auto-generated ID.

---

### Update a movie
```
PUT /movies/{id}
```
Replaces the movie with the given ID. Request body format is the same as creation.

---

### Delete a movie
```
DELETE /movies/{id}
```
Removes the movie with the given ID from the list.

## Data Model

Each movie record contains the following fields:

| Field | Type | Description |
|---|---|---|
| `id` | string | Auto-generated unique identifier |
| `isbn` | string | Book/film ISBN or reference number |
| `title` | string | Title of the movie |
| `director` | object | Director's `firstname` and `lastname` |

## How It Works

The API stores movie data in an in-memory slice (`movies []Movie`). Routes are registered using Gorilla Mux, with a dedicated handler function for each CRUD operation. Since data is held in memory, it resets each time the server restarts.

## Project Structure
```
movies-crud/
├── main.go     # Application entry point, models, and route handlers
├── go.mod      # Module definition and dependencies
└── go.sum      # Dependency checksums
```

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/gorilla/mux` | HTTP request routing |
