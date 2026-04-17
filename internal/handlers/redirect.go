package handlers

import (
    "net/http"
    "strings"
    "time"
    
    "github.com/ZhiRafik/UrlShortenizer/internal/domain"
    "github.com/ZhiRafik/UrlShortenizer/internal/storage"
)

type RedirectHandler struct {
    storage storage.Storage
}

func NewRedirectHandler(s storage.Storage) *RedirectHandler {
    return &RedirectHandler{storage: s}
}

func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    shortCode := strings.TrimPrefix(r.URL.Path, "/")
    if shortCode == "" {
        http.NotFound(w, r)
        return
    }
    
    link, err := h.storage.GetLink(r.Context(), shortCode)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
    
    if link == nil {
        http.NotFound(w, r)
        return
    }
    
    if !link.ExpiresAt.IsZero() && link.ExpiresAt.Before(time.Now()) {
        http.Error(w, "Link expired", http.StatusGone)
        return
    }
    
    click := &domain.ClickStat{
        ShortCode: shortCode,
        Timestamp: time.Now(),
        UserAgent: r.UserAgent(),
        IP:        r.RemoteAddr,
    }
    _ = h.storage.SaveClick(r.Context(), click)
    
    http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}