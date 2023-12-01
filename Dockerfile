# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /

# Install GCC (GNU Compiler Collection)
RUN apt-get update && apt-get install -y gcc

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE $PORT

# Command to run the executable
CMD ["./main"]
