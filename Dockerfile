# Этап сборки
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd/server

# Финальный этап
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]