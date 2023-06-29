# Use the official Go image as the base image
FROM golang:1.17-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a minimal Alpine image as the base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/app .

# Expose the port that the server listens on
EXPOSE 8080

# Set the command to run the Go application
CMD ["./app"]
