# syntax=docker/dockerfile:1

# ========== Stage 1: build ==========
FROM golang:1.25.0 AS builder

WORKDIR /app

# кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# ставим air (для dev)
RUN go install github.com/air-verse/air@latest

# копируем весь проект
COPY . .

# собираем оба бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/bot ./cmd/bot

# ========== Stage 2: runtime ==========
FROM alpine:3.19 AS runtime

WORKDIR /app

COPY --from=builder /bin/api /bin/api
COPY --from=builder /bin/bot /bin/bot

# ========== Stage 3: dev (с air) ==========
FROM builder AS dev

WORKDIR /app
# air уже установлен в /go/bin/air
