package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-url-shortener/internal/handler"
	"go-url-shortener/internal/middleware"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/service"
)

func main() {

	repo := repository.NewMemoryRepository()
	service := service.NewURLService(repo)
	handler := handler.NewURLHandler(service)

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	router.HandleFunc("/{code}", handler.Redirect).Methods("GET")

	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
