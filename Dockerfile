
# Build Stage
FROM golang:1.22 AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary for the API
RUN go build -o ./tmp/main ./cmd/api/main.go

# Final Stage
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/tmp/main /main

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/main"]
