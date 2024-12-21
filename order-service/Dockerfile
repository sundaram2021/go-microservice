FROM ubuntu:22.04 as build

WORKDIR /app

RUN apt-get update && \
    apt-get install -y wget git curl build-essential && \
    wget https://go.dev/dl/go1.22.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz && \
    rm go1.22.6.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service ./cmd/main.go

FROM ubuntu:22.04

WORKDIR /app

COPY --from=build /app/order-service /app/order-service

EXPOSE 8083

ENV DB_HOST=db_host
ENV DB_USER=db_user
ENV DB_PASSWORD=db_password
ENV DB_NAME=db_name
ENV DB_PORT=5432
ENV SSL_MODE=require

CMD ["./order-service"]
