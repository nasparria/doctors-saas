# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest  

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the config directory
COPY --from=builder /app/config ./config

# Copy the wait-for-it script
COPY wait-for-it.sh .

# Make the script executable
RUN chmod +x wait-for-it.sh

# Command to run the executable
CMD ["./main"]