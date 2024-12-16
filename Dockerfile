FROM golang:1.22.2 AS builder

WORKDIR /app

COPY . .

RUN go build -o hello-world-server .

EXPOSE 8080

CMD ["./hello-world-server"]
