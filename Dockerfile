# Start from a Golang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Install PostgreSQL client library
RUN apk add --no-cache postgresql-client

# Build the Go application
RUN go build -o file-processor ./cmd/main.go

# Expose port 8080 (you can change this if necessary)
EXPOSE 8080

# Set the command to run the binary executable
CMD ["./file-processor", "/app/data/txns.csv", "recipient@example.com"]
