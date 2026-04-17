package main

import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/ZhiRafik/UrlShortenizer/internal/storage"
    "github.com/ZhiRafik/UrlShortenizer/internal/handlers"
    "github.com/ZhiRafik/UrlShortenizer/internal/middleware"
)

func main() {
    store := storage.NewMemoryStorage()
    
    shortenHandler := handlers.NewShortenHandler(store)
    redirectHandler := handlers.NewRedirectHandler(store)
    statsHandler := handlers.NewStatsHandler(store)
    
    mux := http.NewServeMux()
    mux.HandleFunc("POST /shorten", shortenHandler.ServeHTTP)
    mux.HandleFunc("GET /stats/{code}", statsHandler.ServeHTTP)
    mux.HandleFunc("GET /{code}", redirectHandler.ServeHTTP)
    
    handler := middleware.Logging(mux)
    handler = middleware.Recovery(handler)
    
    fmt.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}