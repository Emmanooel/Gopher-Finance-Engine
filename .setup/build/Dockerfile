# ---------- BUILD ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia dependências primeiro (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Build do binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# ---------- RUNTIME ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copia apenas o binário
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]