# Go API Boilerplate

This is a boilerplate project for a Go REST API that  includes user authentication with JWT

## Requirements

- Docker
- Docker Compose
- Go 1.22.4+

## Build and Run with Docker Compose

```
docker-compose up 
```

## Access the API

The API server will be available at <http://localhost:3000>.

## API Endpoints

### `POST /login`

**Description:**

- Authenticates a user and returns a JWT token if the login is successful.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}

```

**Response:**

- **Success: `200 OK`**

```json
{
  "jwt": "your-jwt-token"
}
```

- **Error: `400 Bad Request` / `401 Unauthorized` / `500 Internal Server Error`**

### `POST /register`

**Description:**

- Registers a new user and returns a JWT token upon successful registration.

**Request Body:**

```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

**Response:**

- **Success: `200 OK`**

```json
{
  "jwt": "your-jwt-token"
}
```

- **Error: `400 Bad Request` / `500 Internal Server Error`**

### `GET /protected/{id}`

**Description:**

- Accesses a protected route. Requires a valid JWT token in the `Authorization` header.

**Headers:**

```
Authorization: your-jwt-token
```

**URL Parameters:**

- `id` - The ID of the user to access the protected resource.

**Response:**

- **Success: `200 OK`**

```

Access granted to user {id} with email {email}
```

- **Error: `401 Unauthorized`**

## Project Configuration

The configuration for the database and server is managed using environment variables. The `.env` file contains these variables, and they are loaded into the application at runtime.

## MakeFile

The `Makefile` includes the following commands:

>`build`: Builds the Go application
>`run`: Builds and runs the Go application

## Dependencies

- `github.com/joho/godotenv`: For loading environment variables from a .env file
- `github.com/golang-jwt/jwt/v5`: For JWT token generation and validation
- `github.com/lib/pq`: PostgreSQL driver for Go
