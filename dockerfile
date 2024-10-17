# Use the official Golang image with version 1.23.1 as the base image
FROM golang:1.23.1

# Set the working directory inside the container to /app
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .


# Expose port 8099 to the host
EXPOSE 8099

# Run the Go application
CMD ["go", "run", "."]