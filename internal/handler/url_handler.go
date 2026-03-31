package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/dto"
	"github.com/SamuelGuimaraesCosta/URL-Shortener-API/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	var req dto.CreateURLRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	url, err := h.service.CreateShortURL(ctx, req.URL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.CreateURLResponse{
		Code:  url.Code,
		Short: "http://localhost:8080/" + url.Code,
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	params := mux.Vars(r)

	code := params["code"]

	url, err := h.service.ResolveURL(ctx, code)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
