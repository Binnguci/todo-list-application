FROM ubuntu:latest
LABEL authors="binnguci"
FROM golang:1.23.0 AS builder

WORKDIR /app

# Copy go mod and sum files first for better cache utilization
COPY go.mod go.sum ./
RUN go mod download

# Copy all files and build
COPY . .
RUN go build -o todo-app

# Slim Debian for final image
FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /app/todo-app .

EXPOSE 8080

# Command to run the application
CMD ["./todo-app"]
