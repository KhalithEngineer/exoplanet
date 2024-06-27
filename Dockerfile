# Use an official Golang runtime as a parent image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
