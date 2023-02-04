# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

WORKDIR /app

RUN go get github.com/gin-gonic/gin

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]

