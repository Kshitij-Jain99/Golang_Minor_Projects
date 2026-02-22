# go-server

A basic web server written in Go that serves static files and handles simple form submissions and routes. This project demonstrates how to build a minimal HTTP server with custom request handlers using Go's standard `net/http` package.

## Features

- Serves static files from a `static` directory
- Handles form POST requests and returns submitted data
- Includes a sample greeting route at `/hello`
- Built entirely with the Go standard library — no external dependencies

## Requirements

- Go 1.18 or later

## Installation

1. **Clone the repository:**
```bash
   git clone https://github.com/Kshitij-Jain99/Golang_Minor_Projects.git
   cd Golang_Minor_Projects/go-server
```

2. **Run the server:**
```bash
   go run main.go
```

   The server will start listening on **port 8080** by default.

## Usage

### Static Files

Place your static assets (HTML, CSS, JS, images) in the `static` folder. Once the server is running, visit:
```
http://localhost:8080/
```

### Hello Route
```
GET /hello
```

Returns a simple greeting:
```
hello!
```

### Form Submission
```
POST /form
```

The server parses and returns values submitted via an HTML form. Example form:
```html
<form action="/form" method="POST">
  <input type="text" name="name" placeholder="Name">
  <input type="text" name="address" placeholder="Address">
  <button type="submit">Submit</button>
</form>
```

The server will read the `name` and `address` fields and return their values in the response.

## How It Works

The server uses Go's `net/http` package to:

1. Serve static content from `./static` using `http.FileServer`
2. Handle GET and POST requests via `http.Handle` and `http.HandleFunc`
3. Parse submitted form values using `r.ParseForm()` and `r.FormValue()`
4. Listen for incoming connections on port 8080

## Project Structure
```
go-server/
├── static/         # Static assets (HTML, CSS, JS, images)
├── main.go         # Application entry point and route handlers
├── go.mod
└── go.sum
```

## Dependencies

This project uses only the Go standard library — no external modules required.

| Package | Usage |
|---|---|
| `net/http` | HTTP server, routing, and request handling |
| `fmt` | Formatted response output |
| `log` | Error logging |
