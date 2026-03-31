package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/handler"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/middleware"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/repository"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/service"
)

func main() {

	repo := repository.NewMemoryRepository()

	urlService := service.NewURLService(repo)

	urlHandler := handler.NewURLHandler(urlService)

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimit)

	router.HandleFunc("/shorten", urlHandler.ShortenURL).Methods("POST")
	router.HandleFunc("/{code}", urlHandler.Redirect).Methods("GET")

	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
