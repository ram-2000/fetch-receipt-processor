# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o receipt-service .

# Stage 2: Create a minimal runtime image
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/receipt-service .

# Expose the port your application runs on
EXPOSE 8080

# Run the Go application
CMD ["./receipt-service"]
