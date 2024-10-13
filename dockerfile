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

# Set environment variables for the Twitter API keys and tokens
ENV GOTWI_API_KEY="Qfl4sHzT8BCfPmf0YkrTl3cHj"
ENV GOTWI_API_KEY_SECRET="IpzNURsHZLzw0kGKdctbv6jnOva1nIJqDAciG6TxXMkym9K7AU"
ENV ACCESS_TOKEN="1559069157916024834-GGRAm1LMd0Hy0RnQXDZm9Cdq3qaY6m"
ENV ACCESS_TOKEN_SECRET="VAwFSu7NopZxn6aXpoaD76Ja7rsA0fKoKwB3tagY8Zflz"

# Expose port 8099 to the host
EXPOSE 8099

# Run the Go application
CMD ["go", "run", "."]