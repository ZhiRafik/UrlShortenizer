package storage

import (
    "context"
    "github.com/ZhiRafik/UrlShortenizer/internal/domain"
)

type Storage interface {
    SaveLink(ctx context.Context, link *domain.Link) error
    GetLink(ctx context.Context, shortCode string) (*domain.Link, error)
    DeleteLink(ctx context.Context, shortCode string) error
    SaveClick(ctx context.Context, click *domain.ClickStat) error
    GetStats(ctx context.Context, shortCode string) (*domain.StatsResponse, error)
    ListExpired(ctx context.Context) ([]*domain.Link, error)
}