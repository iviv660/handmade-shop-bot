FROM --platform=$BUILDPLATFORM golang:1.25.1-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /app ./cmd/app && chmod +x /app

FROM --platform=$TARGETPLATFORM alpine:3.20
WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
