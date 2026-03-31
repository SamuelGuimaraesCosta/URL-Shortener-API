package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/model"
)

type MemoryRepository struct {
	data map[string]*model.URL
	mu   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]*model.URL),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, url *model.URL) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[url.Code] = url

	return nil
}

func (r *MemoryRepository) Get(ctx context.Context, code string) (*model.URL, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.data[code]

	if !ok {
		return nil, errors.New("url not found")
	}

	return url, nil
}

func (r *MemoryRepository) IncrementHits(ctx context.Context, code string) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	if url, ok := r.data[code]; ok {
		url.Hits++
		return nil
	}

	return errors.New("url not found")
}
