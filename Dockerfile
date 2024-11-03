# Build stage
FROM golang:1.21-alpine AS builder

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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Final stage
FROM alpine:3.18

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates curl tzdata

# Create non-root user
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/main .

# Copy migrations
COPY --from=builder /build/migrations ./migrations

# Change ownership of the application directory
RUN chown -R appuser:appuser /app

# Use non-root user
USER appuser

# Command to run the executable
CMD ["./main"]