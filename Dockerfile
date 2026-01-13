# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o aggregator .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/aggregator .

# Create results directory
RUN mkdir -p /app/results

# Set default entrypoint
ENTRYPOINT ["./aggregator"]

# Default command (can be overridden)
CMD ["--help"]
