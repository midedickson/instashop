# Build stage
FROM golang:1.22-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git curl

# Set working directory
WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Final stage
FROM debian:bullseye-slim

# Install necessary runtime dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN useradd -m appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/main .

# Copy env file from builder
COPY --from=builder /build/.env .

# # Copy migrations
# COPY --from=builder /build/migrations ./migrations

# Change ownership of the application directory
RUN chown -R appuser:appuser /app

# Use non-root user
USER appuser

EXPOSE 8080

# Command to run the executable
CMD ["./main"]