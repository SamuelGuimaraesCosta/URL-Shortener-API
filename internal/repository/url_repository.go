package repository

import (
	"context"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/model"
)

type URLRepository interface {
	Save(ctx context.Context, url *model.URL) error
	Get(ctx context.Context, code string) (*model.URL, error)
	IncrementHits(ctx context.Context, code string) error
}
