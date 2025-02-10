# Specify the Go version
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} AS builder

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main ./cmd/main

# Create a minimal final image
FROM alpine:latest

# Install necessary certificates for secure connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built application binary from the builder stage
COPY --from=builder /app/main .

# Expose the service port
EXPOSE 8080

# Command to run the app
CMD ["./main"]
