# Stage 1: Builder
FROM golang:1.25.1-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Собираем бинарь под ту же архитектуру, что и образ (без TARGETARCH)
RUN CGO_ENABLED=0 go build -o /app ./cmd/app && chmod +x /app

# Stage 2: Runner
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
