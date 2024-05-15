FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 9000
ENTRYPOINT ["go", "run", "main.go"]