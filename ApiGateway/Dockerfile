FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 44333
ENTRYPOINT ["go", "run", "main.go"]