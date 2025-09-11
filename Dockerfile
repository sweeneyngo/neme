# --- Build stage ---
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install git for modules
RUN apk add --no-cache git

# Copy modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy everything
COPY . .

RUN go build -o neme ./cmd/web

# --- Final stage ---
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/neme ./

COPY .cache/cache-order2.gob .cache/cache-order2.gob

EXPOSE 8080

# Run the web server
CMD ["./neme"]
