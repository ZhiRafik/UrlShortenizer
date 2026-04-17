package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/ZhiRafik/UrlShortenizer/internal/domain"
    "github.com/ZhiRafik/UrlShortenizer/internal/storage"
    "github.com/ZhiRafik/UrlShortenizer/pkg/utils"
)

type ShortenHandler struct {
    storage storage.Storage
}

func NewShortenHandler(s storage.Storage) *ShortenHandler {
    return &ShortenHandler{storage: s}
}

type shortenRequest struct {
    URL       string `json:"url"`
    ExpiresIn int    `json:"expires_in,omitempty"`
}

type shortenResponse struct {
    ShortURL  string `json:"short_url"`
    ShortCode string `json:"short_code"`
}

func (h *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var req shortenRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    if req.URL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    
    shortCode := utils.GenerateShortCode()
    
    link := &domain.Link{
        ShortCode:   shortCode,
        OriginalURL: req.URL,
        CreatedAt:   time.Now(),
        Clicks:      0,
    }
    
    if req.ExpiresIn > 0 {
        link.ExpiresAt = time.Now().Add(time.Duration(req.ExpiresIn) * time.Hour)
    }
    
    if err := h.storage.SaveLink(r.Context(), link); err != nil {
        http.Error(w, "Failed to save link", http.StatusInternalServerError)
        return
    }
    
    resp := shortenResponse{
        ShortURL:  "http://localhost:8080/" + shortCode,
        ShortCode: shortCode,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(resp)
}