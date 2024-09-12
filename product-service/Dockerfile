# Stage 1: Build the Go application
FROM ubuntu:22.04 as build

WORKDIR /app

# Install dependencies to build the Go app
RUN apt-get update && \
    apt-get install -y wget git curl build-essential && \
    wget https://go.dev/dl/go1.22.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz && \
    rm go1.22.6.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o product-service ./cmd/main.go

# Stage 2: Create a minimal image with only the binary
FROM ubuntu:22.04

WORKDIR /app

# Copy the built binary from the first stage
COPY --from=build /app/product-service /app/product-service

# Expose the port
EXPOSE 8082

ENV DB_HOST=db_host
ENV DB_USER=db_user
ENV DB_PASSWORD=db_password
ENV DB_NAME=db_name
ENV DB_PORT=5432
ENV SSL_MODE=require

# Start the Go application
CMD ["./product-service"]
