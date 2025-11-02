# Build stage
FROM golang:1.25-rc AS builder

# Set working directory
WORKDIR /app

# Copy all files and directories from the build context
COPY . .

# Download dependencies
RUN go mod download

# Build the binary with CGO disabled for static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o serverinfo main.go

# Runtime stage
FROM debian:bullseye

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/serverinfo .

# Ensure the binary is executable
RUN chmod +x ./serverinfo

# Set entrypoint to the binary
ENTRYPOINT ["./serverinfo"]

# Default command (optional, can be overridden by user)
CMD ["help"]