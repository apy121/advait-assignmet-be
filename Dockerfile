# Specify the Go version
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} AS builder

WORKDIR /app

# Copy project files
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env /app/.env

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main

# Final Stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Execute the binary
ENTRYPOINT ["./main"]