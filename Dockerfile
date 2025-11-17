
FROM golang:1.24-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download

COPY . .git


RUN go build -o /usr/local/bin/todo-app ./...

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /usr/local/bin/todo-app .


EXPOSE 8080

CMD ["./todo-app"]
