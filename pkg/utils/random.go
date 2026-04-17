package utils

import (
    "math/rand"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GenerateShortCode() string {
    code := make([]byte, 6)
    for i := range code {
        code[i] = charset[rand.Intn(len(charset))]
    }
    return string(code)
}