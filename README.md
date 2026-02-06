# Go API

A simple REST API built with Go that provides user authentication and coin balance management.

## Features

- User authentication with token-based authorization
- Get user coin balance endpoint
- Mock database for testing
- Middleware for request authorization

## Prerequisites

- Go 1.25.6 or higher
- Postman (or any API testing tool)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/FreshPrince99/Go_API.git
cd Go_API
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8000`

## API Endpoints

### Get Coin Balance

Get the coin balance for a specific user.

**Endpoint:** `GET /account/coins`

**Query Parameters:**
- `username` (required): The username of the account
- `Authorization` (required): The authentication token

**Example Request:**
```
GET http://localhost:8000/account/coins/?username=alex&Authorization=123ABC
```

**Example Response:**
```json
{
  "code": 200,
  "balance": 100
}
```

## Mock Users

The application includes three mock users for testing:

| Username | Token    | Coins |
|----------|----------|-------|
| alex     | 123ABC   | 100   |
| john     | 456DEF   | 200   |
| marie    | 789GHI   | 300   |

## Project Structure

```
Go_API/
├── api/              # API models and response handlers
├── cmd/api/          # Application entry point
├── internal/
│   ├── handlers/     # HTTP request handlers
│   ├── middleware/   # HTTP middleware (authorization)
│   └── tools/        # Database interfaces and mock data
├── go.mod
└── README.md
```

## Technologies Used

- [Go Chi](https://github.com/go-chi/chi) - Lightweight router
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger
- [Gorilla Schema](https://github.com/gorilla/schema) - Query parameter decoder

## License

MIT
