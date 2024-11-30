# Use the latest Go image (Go 1.21 as of now)
FROM golang:1.23.3-alpine

# Set the working directory
WORKDIR /app

# Copy Go module files (just go.mod)
COPY go.mod ./

# Download dependencies (go.sum will be ignored if it doesn't exist)
RUN go mod download

# Copy the source code
COPY . .

# Run tests (this will run `go test` before building the app)
RUN go test -v ./...

# Build the Go application
RUN go build -o app .

# Expose the port
EXPOSE 8080

# Run the Go application
CMD ["./app"]