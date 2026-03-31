package service

import (
	"context"
	"testing"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/repository"
)

func TestCreateShortURL(t *testing.T) {

	repo := repository.NewMemoryRepository()

	service := NewURLService(repo)

	ctx := context.Background()

	url, err := service.CreateShortURL(ctx, "https://google.com")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if url.Code == "" {
		t.Fatalf("expected generated code")
	}

	if url.OriginalURL != "https://google.com" {
		t.Fatalf("original URL mismatch")
	}
}

func TestResolveURL(t *testing.T) {

	repo := repository.NewMemoryRepository()

	service := NewURLService(repo)

	ctx := context.Background()

	created, err := service.CreateShortURL(ctx, "https://google.com")

	if err != nil {
		t.Fatalf("error creating url: %v", err)
	}

	resolved, err := service.ResolveURL(ctx, created.Code)

	if err != nil {
		t.Fatalf("error resolving url")
	}

	if resolved.OriginalURL != "https://google.com" {
		t.Fatalf("expected original url")
	}
}
