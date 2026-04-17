package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    
    "github.com/ZhiRafik/UrlShortenizer/internal/storage"
)

type StatsHandler struct {
    storage storage.Storage
}

func NewStatsHandler(s storage.Storage) *StatsHandler {
    return &StatsHandler{storage: s}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    path := strings.TrimPrefix(r.URL.Path, "/stats/")
    if path == "" {
        http.Error(w, "Short code required", http.StatusBadRequest)
        return
    }
    
    stats, err := h.storage.GetStats(r.Context(), path)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
    
    if stats == nil {
        http.NotFound(w, r)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}