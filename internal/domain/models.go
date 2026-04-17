package domain

import "time"

type Link struct {
    ShortCode   string    `json:"short_code"`
    OriginalURL string    `json:"original_url"`
    CreatedAt   time.Time `json:"created_at"`
    ExpiresAt   time.Time `json:"expires_at,omitempty"`
    Clicks      int       `json:"clicks"`
}

type ClickStat struct {
    ShortCode string    `json:"short_code"`
    Timestamp time.Time `json:"timestamp"`
    UserAgent string    `json:"user_agent"`
    IP        string    `json:"ip"`
}

type StatsResponse struct {
    ShortCode   string      `json:"short_code"`
    OriginalURL string      `json:"original_url"`
    TotalClicks int         `json:"total_clicks"`
    Clicks      []ClickStat `json:"clicks,omitempty"`
}