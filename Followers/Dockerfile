FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8060
ENTRYPOINT ["go", "run", "main.go"]