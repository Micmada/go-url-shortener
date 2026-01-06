---
description: High-performance URL shortener built with Go, Gin, and GORM with JWT authentication
details: >
  A fast, lightweight URL shortening service built in Go using the Gin web
  framework and GORM ORM for database operations. Features secure user
  authentication with JWT tokens, bcrypt password hashing, and RESTful API
  endpoints for URL shortening and redirection. Users can register accounts,
  log in to receive authentication tokens, and manage shortened URLs through
  a clean API interface. The application uses SQLite for efficient local storage,
  supports automatic database migrations via GORM, and handles URL validation
  and unique short code generation. Built with Go's concurrency model for high
  performance, minimal resource usage, and fast request handling. The modular
  architecture separates routing, authentication middleware, and database models,
  making it easy to extend with additional features or switch to PostgreSQL or
  MySQL for production deployments.
technologies:
  - go
  - gin
  - gorm
  - jwt
  - sqlite
  - bcrypt
hostedUrl: 
---



# Go URL Shortener

A simple URL shortener built with Go, Gin, and GORM.

## Features
- Shorten URLs
- Redirect short URLs to original URLs
- User authentication (registration & login)
- JWT-based authentication

## Tech Stack
- **Go** (Golang)
- **Gin** (HTTP Web Framework)
- **GORM** (ORM for Go)
- **SQLite** (Database)
- **bcrypt** (Password hashing)
- **JWT** (Authentication)

## Installation

### Prerequisites
- Go 1.18+

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/Micmada/go-url-shortener.git
   cd go-url-shortener
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

### URL Shortening
- **POST /api/shorten** – Shortens a URL
- **GET /api/:shortURL** – Redirects to the original URL

### User Authentication
- **POST /api/register** – Registers a new user
- **POST /api/login** – Logs in a user and returns a JWT token

## Configuration
- The server runs on **port 8080** by default.
- The database file is **shortener.db** (SQLite).


Feel free to contribute and improve this project!

