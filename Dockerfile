# Build stage
FROM golang:1.25.4-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Install swag for Swagger docs generation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY web/ ./web/

# Generate Swagger docs and build
RUN /go/bin/swag init -g cmd/server/main.go -o docs
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o taskwarrior-api ./cmd/server

# Using archlinux image to install Taskwarrior 3.x
FROM archlinux/archlinux:latest

# Install Taskwarrior 3.x and dependencies
RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm \
    task \
    ca-certificates \
    wget \
    && pacman -Scc --noconfirm

# Set working directory
WORKDIR /app

# Copy API binary from Go builder
COPY --from=builder /build/taskwarrior-api .

# Copy web templates
COPY --from=builder /build/web ./web

# Copy entrypoint script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Create a default non-root user (UID/GID will be modified at runtime)
RUN groupadd -g 1000 appuser && \
    useradd -u 1000 -g 1000 -m -s /bin/bash appuser

# Create directory for Taskwarrior data and initialize config
RUN mkdir -p /home/appuser/.task && \
    echo "data.location=/home/appuser/.task" > /home/appuser/.taskrc && \
    echo "confirmation=no" >> /home/appuser/.taskrc && \
    chown -R appuser:appuser /home/appuser/.task /home/appuser/.taskrc

# Change ownership of app directory
RUN chown -R appuser:appuser /app

# Expose port
EXPOSE 8080

# Set environment variables with defaults
ENV PUID=1000
ENV PGID=1000

LABEL org.opencontainers.image.source "https://github.com/dotbinio/taskwarrior-api"

ENTRYPOINT ["/app/entrypoint.sh"]
