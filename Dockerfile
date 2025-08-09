# Dockerfile
# Build
FROM golang:1.24.0-alpine3.20 as builder
# Install dependencies for builds
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

COPY config/config.yaml ./config/config.yaml
# Copy go.mod and go.sum first (for caching modules)
COPY go.mod go.sum ./
# Download Go modules
RUN go mod download
# Copy the rest of the code
COPY . .
# Build the Go binary
RUN go build -o message-service ./cmd/main.go



# Run
FROM alpine:latest
# Create a non-root user for security
RUN adduser -D appuser
# Set working directory
WORKDIR /app
# Copy binary from builder
COPY --from=builder /app/message-service .
# Copy config file
COPY --from=builder /app/config/config.yaml ./config/config.yaml


# Change to non-root user
USER appuser

# Run the service
CMD ["./message-service"]