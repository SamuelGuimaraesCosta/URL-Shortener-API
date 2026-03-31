# URL-Shortener-API

A production-style URL shortener API built with **Go (Golang)**.
This project demonstrates a layered backend architecture commonly used in professional Go services.

The API allows users to shorten URLs and redirect to the original links while tracking access statistics.

---

## Repository

https://github.com/SamuelGuimaraesCosta/URL-Shortener-API

---

## Features

* Create shortened URLs
* Redirect shortened URLs to the original link
* Track access count (hits)
* Thread-safe in-memory storage
* Layered architecture (Handler → Service → Repository)
* Middleware logging
* Cryptographically secure short code generation

---

## Tech Stack

* **Go**
* **Gorilla Mux (HTTP Router)**
* **REST API**
* **Mutex for concurrency safety**
* **Standard Library**

---

## Project Architecture

The project follows a clean backend architecture structure:

```
URL-Shortener-API
│
├── cmd
│   └── server
│       └── main.go
│
├── internal
│   ├── handler
│   ├── service
│   ├── repository
│   ├── middleware
│   └── model
│
├── pkg
│   └── utils
│
└── go.mod
```

### Layers

**Handler**

* HTTP endpoints
* request/response handling

**Service**

* business logic
* URL creation and retrieval

**Repository**

* data storage abstraction

**Middleware**

* logging and request lifecycle management

---

## Installation

Clone the repository:

```
git clone https://github.com/SamuelGuimaraesCosta/URL-Shortener-API.git
```

Enter the project folder:

```
cd URL-Shortener-API
```

Install dependencies:

```
go mod tidy
```

Run the server:

```
go run ./cmd/server
```

The API will start at:

```
http://localhost:8080
```

---

## API Endpoints

### Create Short URL

POST `/shorten`

Request:

```
{
  "url": "https://google.com"
}
```

Response:

```
{
  "short": "http://localhost:8080/Ab3KpQ"
}
```

---

### Redirect to Original URL

GET `/{code}`

Example:

```
http://localhost:8080/Ab3KpQ
```

The server will redirect to the stored original URL.

---

## Example using curl

Create a short URL:

```
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://google.com"}'
```

---

## Future Improvements

* PostgreSQL storage
* Redis caching
* Rate limiting
* URL expiration
* Authentication (JWT)
* Metrics and monitoring (Prometheus)
* Docker container support
* Unit and integration tests

---

## License

MIT
