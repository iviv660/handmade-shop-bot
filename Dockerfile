FROM golang:1.25.1-bookworm AS builder
WORKDIR /src
RUN apt-get update && apt-get install -y --no-install-recommends git ca-certificates file \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -trimpath -ldflags "-s -w" -o /app ./cmd/app && chmod +x /app

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=builder /app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
