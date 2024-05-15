FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8093
ENTRYPOINT ["go", "run", "main.go"]