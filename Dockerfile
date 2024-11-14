# Build stage
FROM golang:1.22-alpine AS builder

# Install git and build dependencies
RUN apk add --no-cache git build-base

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# TODO: Need to invoke 'make' command to build the Go binaries
# Build the application
# Adding more flags for complete static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o main .

# Final stage
FROM scratch

# Copy ca-certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/main /main

# Copy any additional config files if needed
# COPY --from=builder /app/config /config

# Expose port (adjust as needed)
EXPOSE 8080

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    # non-HTTP check:
    # CMD ["/main", "health"]
    # HTTP check: (use one of these)
    CMD ["/main", "-http-get=http://localhost:8080/health"]

# Command to run
ENTRYPOINT ["/main"] 