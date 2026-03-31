package service

import (
	"context"
	"errors"
	"time"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/model"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/repository"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/pkg/utils"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortURL(ctx context.Context, original string) (*model.URL, error) {

	if original == "" {
		return nil, errors.New("url cannot be empty")
	}

	code := utils.GenerateCode()

	url := &model.URL{
		Code:        code,
		OriginalURL: original,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Save(ctx, url)

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) ResolveURL(ctx context.Context, code string) (*model.URL, error) {

	url, err := s.repo.Get(ctx, code)

	if err != nil {
		return nil, err
	}

	if url.IsExpired() {
		return nil, errors.New("url expired")
	}

	_ = s.repo.IncrementHits(ctx, code)

	return url, nil
}
