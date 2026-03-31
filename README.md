Go URL Shortener

A simple URL shortener API built with Go.

This project demonstrates a minimal REST API using the Go standard
library.

Features

-   Create shortened URLs
-   Redirect to original URL
-   Thread-safe in-memory storage
-   No external dependencies

Tech Stack

-   Go
-   net/http
-   JSON API

Project Structure

main.go -> server entry point handlers.go -> API handlers store.go ->
in-memory storage

Running the Project

1.  Clone the repository

git clone https://github.com/SamuelGuimaraesCosta/URL-Shortener-API

2.  Run the application

go run .

Server will start on:

http://localhost:8080

API Usage

Shorten URL

POST /shorten

Request:

{ “url”: “https://google.com” }

Response:

{ “short”: “http://localhost:8080/abc123” }

Redirect

Access:

http://localhost:8080/abc123

The service will redirect to the original URL.

Possible Improvements

-   Persistent storage (PostgreSQL / Redis)
-   Custom short codes
-   Expiration time for links
-   Docker support
-   Rate limiting
-   Metrics and logging

License

MIT
