# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the Go server
EXPOSE 8080

# Start the application
CMD ["./main"]
