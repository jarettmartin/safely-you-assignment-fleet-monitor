# =====================================
# 1. Build stage
# =====================================
FROM golang:1.25-alpine AS builder

# Create and set working directory
WORKDIR /app

# Cache dependencies first (faster rebuilds)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary to a known location
RUN go build -o /fleet-monitor .

# =====================================
# 2. Runtime stage
# =====================================
FROM alpine:latest

WORKDIR /app

# Copy built binary and any required runtime assets
COPY --from=builder /fleet-monitor /app/fleet-monitor
COPY devices.csv .
COPY .env .

# Expose API port
EXPOSE 6733

# Use absolute path so it works regardless of host directory name
CMD ["/app/fleet-monitor"]
