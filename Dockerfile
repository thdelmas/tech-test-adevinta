# Use golang base image from Docker Hub
FROM golang:1.22.2-alpine AS builder

# Install git
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies that are available publicly
RUN go mod download

# Install mockgen
RUN go install github.com/golang/mock/mockgen@latest

# Copy source code
COPY . .

# Create mocks directory and generate mocks
RUN mkdir -p mocks
RUN mockgen -source=services/fizzbuzz.go -destination=mocks/mock_fizzbuzz.go -package=mocks

RUN go mod tidy

# Run go generate to ensure all other mocks are generated
RUN go generate ./...

# Run tests
RUN go test ./... -v &> test-results.log || (echo "TESTS FAILED" && cat test-results.log && exit 1)

# Build the Go app
RUN go build -o fizzbuzz-api .

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/fizzbuzz-api .
COPY --from=builder /app/test-results.log .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./fizzbuzz-api"]