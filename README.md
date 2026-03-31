# URL-Shortener-API

A production-style **URL Shortener REST API** built with **Go (Golang)** demonstrating a layered backend architecture, middleware, testing, and containerization.

This project was designed to simulate a **real backend service structure** used in modern Go applications.

Repository:
https://github.com/SamuelGuimaraesCosta/URL-Shortener-API

---

# Features

* URL shortening service
* HTTP redirect to original URL
* Thread-safe in-memory storage
* Hit counter tracking
* Rate limiting middleware
* Structured logging middleware
* Layered architecture (Handler → Service → Repository)
* OpenAPI / Swagger documentation
* Docker container support
* Unit tests

---

# Tech Stack

* Go
* Gorilla Mux Router
* REST API
* Docker
* OpenAPI (Swagger)
* Unit Testing

---

# Project Structure

```
URL-Shortener-API
│
├── api
│   └── openapi.yaml
│
├── cmd
│   └── server
│       └── main.go
│
├── internal
│   ├── dto
│   │   └── url_dto.go
│   │
│   ├── handler
│   │   └── url_handler.go
│   │
│   ├── middleware
│   │   ├── logging.go
│   │   └── rate_limit.go
│   │
│   ├── model
│   │   └── url.go
│   │
│   ├── repository
│   │   ├── memory_repository.go
│   │   └── url_repository.go
│   │
│   └── service
│       ├── url_service.go
│       └── url_service_test.go
│
├── pkg
│   └── utils
│       └── generator.go
│
├── Dockerfile
├── go.mod
└── README.md
```

---

# Running the Project

Install dependencies:

```
go mod tidy
```

Run the API:

```
go run ./cmd/server
```

Server will start at:

```
http://localhost:8080
```

---

# API Endpoints

## Create Short URL

POST `/shorten`

Request

```
{
  "url": "https://google.com"
}
```

Response

```
{
  "code": "Ab3KpQ",
  "short": "http://localhost:8080/Ab3KpQ"
}
```

---

## Redirect

GET

```
/{code}
```

Example

```
http://localhost:8080/Ab3KpQ
```

Returns:

```
HTTP 302 Redirect
```

---

# OpenAPI / Swagger

The API specification is located in:

```
api/openapi.yaml
```

You can visualize the API documentation using **Swagger UI**.

Run Swagger UI with Docker:

```
docker run -p 8081:8080 \
-v $(pwd)/api:/usr/share/nginx/html/api \
swaggerapi/swagger-ui
```

Open in browser:

```
http://localhost:8081
```

---

# Docker

Build the Docker image:

```
docker build -t url-shortener-api .
```

Run the container:

```
docker run -p 8080:8080 url-shortener-api
```

The API will be available at:

```
http://localhost:8080
```

---

# Unit Tests

Run tests with:

```
go test ./...
```

Example test coverage includes:

* URL creation
* URL resolution
* repository integration

Example test file:

```
internal/service/url_service_test.go
```

---

# Example CURL Request

Create a shortened URL:

```
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://google.com"}'
```

---

# Possible Improvements

Future improvements for production use:

* PostgreSQL storage
* Redis cache
* URL expiration
* JWT authentication
* Metrics with Prometheus
* Distributed rate limiting
* Kubernetes deployment

---

# License

MIT
