package storage

import (
    "context"
    "sync"
    "time"
    "github.com/ZhiRafik/UrlShortenizer/internal/domain"
)

type MemoryStorage struct {
    links map[string]*domain.Link
    stats map[string][]domain.ClickStat
    mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        links: make(map[string]*domain.Link),
        stats: make(map[string][]domain.ClickStat),
    }
}

func (s *MemoryStorage) SaveLink(ctx context.Context, link *domain.Link) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.links[link.ShortCode] = link
    return nil
}

func (s *MemoryStorage) GetLink(ctx context.Context, shortCode string) (*domain.Link, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    link, ok := s.links[shortCode]
    if !ok {
        return nil, nil
    }
    return link, nil
}

func (s *MemoryStorage) DeleteLink(ctx context.Context, shortCode string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.links, shortCode)
    return nil
}

func (s *MemoryStorage) SaveClick(ctx context.Context, click *domain.ClickStat) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    s.stats[click.ShortCode] = append(s.stats[click.ShortCode], *click)
    
    if link, ok := s.links[click.ShortCode]; ok {
        link.Clicks++
    }
    return nil
}

func (s *MemoryStorage) GetStats(ctx context.Context, shortCode string) (*domain.StatsResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    link, ok := s.links[shortCode]
    if !ok {
        return nil, nil
    }
    
    clicks := s.stats[shortCode]
    
    return &domain.StatsResponse{
        ShortCode:   shortCode,
        OriginalURL: link.OriginalURL,
        TotalClicks: len(clicks),
        Clicks:      clicks,
    }, nil
}

func (s *MemoryStorage) ListExpired(ctx context.Context) ([]*domain.Link, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    var expired []*domain.Link
    now := time.Now()
    
    for _, link := range s.links {
        if !link.ExpiresAt.IsZero() && link.ExpiresAt.Before(now) {
            expired = append(expired, link)
        }
    }
    return expired, nil
}